package ethereum

import (
	"bytes"
	"context"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/ethereum/client/codec"
)

func (c *Contract) generateTransactionId() uint32 {
	c.transactionsLock.Lock()
	c.transactionsToGrab += 1
	c.transactionsLock.Unlock()

	return c.transactionsToGrab
}

func (c *Client) generateContractId() uint32 {
	c.contractsLock.Lock()
	defer func() {
		c.contractsIdToGrab += 1
		c.contractsLock.Unlock()
	}()

	return c.contractsIdToGrab
}

func (c *Client) getContract(contractId uint32) (*Contract, errno.Error) {
	c.contractsLock.RLock()
	defer c.contractsLock.RUnlock()
	if contract, ok := c.contracts[contractId]; ok {
		return contract, errno.ErrorNone
	}

	return nil, errno.ErrorEthereumContractNotFound
}

func verifyInputs(inputsBytes [][]byte, method *contractMethod) ([]interface{}, errno.Error) {
	if len(inputsBytes) != len(method.inputs) {
		return nil, errno.ErrorEthereumInvalidContractMethodInput
	}

	var inputs []interface{}
	for idx, input := range inputsBytes {
		inputType := method.inputs[idx]
		if inputType == "common.Address" {
			inputs = append(inputs, ethCommon.BytesToAddress(input))
			continue
		}

		decoder, err := codec.Converter(inputType).Decoder()
		if err != nil {
			return nil, errno.ErrorEthereumUnsupportedDataType
		}

		val, err := decoder(input)
		if err != nil {
			return nil, errno.ErrorEthereumParseInputTypeFailed
		}

		inputs = append(inputs, val)
	}

	return inputs, 0
}

func (f *Factory) toEcdsa(module common.Module, privKeyPtr, privKeySize uint32) (*ecdsa.PrivateKey, errno.Error) {
	pkBytes, err0 := f.ReadBytes(module, privKeyPtr, privKeySize)
	if err0 != 0 {
		return nil, err0
	}

	privateKey, err := crypto.ToECDSA(pkBytes)
	if err != nil {
		return nil, errno.ErrorEthereumInvalidPrivateKey
	}

	return privateKey, 0
}

func (f *Factory) W_ethDeployContract(
	ctx context.Context,
	module common.Module,
	clientId,
	chainIdPtr, chainIdSize,
	binPtr, binLen,
	abiPtr, abiSize,
	privKeyPtr, privKeySize,
	addressPtr,
	methodsSizePtr,
	contractIdPtr,
	transactionIdPtr uint32,
) errno.Error {
	client, err0 := f.getClient(clientId)
	if err0 != 0 {
		return err0
	}

	abiJson, err0 := f.ReadBytes(module, abiPtr, abiSize)
	if err0 != 0 {
		return err0
	}

	chainId, err0 := f.ReadBigInt(module, chainIdPtr, chainIdSize)
	if err0 != 0 {
		return err0
	}

	parsedAbi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return errno.ErrorEthereumParsingAbiFailed
	}

	privateKey, err0 := f.toEcdsa(module, privKeyPtr, privKeySize)
	if err0 != 0 {
		return err0
	}

	bin, err0 := f.ReadString(module, binPtr, binLen)
	if err0 != 0 {
		return err0
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return errno.ErrorEthereumBindTransactorFailed
	}

	address, transaction, boundContract, err := bind.DeployContract(auth, parsedAbi, ethCommon.FromHex(bin), client)
	if err != nil {
		return errno.ErrorEthereumDeployFailed
	}

	contract, err0 := f.handleBoundContractSize(module, contractIdPtr, methodsSizePtr, client, parsedAbi, boundContract)
	if err0 != 0 {
		return err0
	}

	tx := &Transaction{
		Transaction: transaction,
		Id:          contract.generateTransactionId(),
	}

	if err0 := f.WriteBytes(module, addressPtr, address[:]); err0 != 0 {
		return err0
	}

	if err0 := f.WriteUint32Le(module, transactionIdPtr, tx.Id); err0 != 0 {
		return err0
	}

	contract.transactionsLock.Lock()
	contract.transactions[tx.Id] = tx
	contract.transactionsLock.Unlock()

	return 0
}

func (f *Factory) handleBoundContractSize(module common.Module, contractIdPtr, methodsSizePtr uint32, client *Client, abi abi.ABI, contract *bind.BoundContract) (*Contract, errno.Error) {
	c := Contract{
		BoundContract: contract,
		clientId:      client.Id,
		Id:            client.generateContractId(),
		methods:       make(map[string]*contractMethod),
		transactions:  make(map[uint32]*Transaction),
	}

	var methodList []string
	for _, method := range abi.Methods {
		methodList = append(methodList, method.Name)

		var inputs []string
		var outputs []string
		for _, input := range method.Inputs {
			inputs = append(inputs, input.Type.GetType().String())
		}

		for _, output := range method.Outputs {
			outputs = append(outputs, output.Type.GetType().String())
		}

		c.methodsLock.Lock()
		c.methods[method.Name] = &contractMethod{
			inputs:   inputs,
			outputs:  outputs,
			constant: method.IsConstant(), // if false then needs to call transaction
		}
		c.methodsLock.Unlock()
	}

	if err := f.WriteStringSliceSize(module, methodsSizePtr, methodList); err != 0 {
		return nil, err
	}

	if err := f.WriteUint32Le(module, contractIdPtr, c.Id); err != 0 {
		return nil, err
	}

	client.contractsLock.Lock()
	client.contracts[c.Id] = &c
	client.contractsLock.Unlock()

	return &c, 0
}

func (f *Factory) W_ethTransactContract(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	chainIdPtr,
	chainIdSize,
	methodPtr,
	methodLen,
	privKeyPtr,
	privKeySize,
	inputsPtr,
	inputsSize,
	transactionIdPtr uint32,
) errno.Error {
	client, err0 := f.getClient(clientId)
	if err0 != 0 {
		return err0
	}

	chainId, err0 := f.ReadBigInt(module, chainIdPtr, chainIdSize)
	if err0 != 0 {
		return err0
	}

	contract, err0 := client.getContract(contractId)
	if err0 != 0 {
		return err0
	}

	methodString, err0 := f.ReadString(module, methodPtr, methodLen)
	if err0 != 0 {
		return err0
	}

	method, ok := contract.methods[methodString]
	if !ok {
		return errno.ErrorEthereumContractMethodNotFound
	}

	if method.constant {
		return errno.ErrorEthereumCannotTransactFreeMethod
	}

	inputsBytes, err0 := f.ReadBytesSlice(module, inputsPtr, inputsSize)
	if err0 != 0 {
		return err0
	}

	inputs, err0 := verifyInputs(inputsBytes, method)
	if err0 != 0 {
		return err0
	}

	privateKey, err0 := f.toEcdsa(module, privKeyPtr, privKeySize)
	if err0 != 0 {
		return err0
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return errno.ErrorEthereumBindTransactorFailed
	}

	transaction, err := contract.Transact(auth, methodString, inputs...)
	if err != nil {
		return errno.ErrorEthereumTransactMethodFailed
	}

	tx := &Transaction{
		Transaction: transaction,
		Id:          contract.generateTransactionId(),
	}

	if err0 := f.WriteUint32Le(module, transactionIdPtr, tx.Id); err0 != 0 {
		return err0
	}

	contract.transactionsLock.Lock()
	contract.transactions[tx.Id] = tx
	contract.transactionsLock.Unlock()

	return 0
}

func (f *Factory) W_ethNewContractSize(
	ctx context.Context,
	module common.Module,
	clientId,
	abiPtr,
	abiSize,
	addressPtr,
	addressLen,
	methodsSizePtr,
	contractPtr uint32,
) errno.Error {
	client, err0 := f.getClient(clientId)
	if err0 != 0 {
		return err0
	}

	abiJson, err0 := f.ReadBytes(module, abiPtr, abiSize)
	if err0 != 0 {
		return err0
	}

	parsedAbi, err := abi.JSON(bytes.NewReader(abiJson))
	if err != nil {
		return errno.ErrorEthereumParsingAbiFailed
	}

	var contractAddress ethCommon.Address
	address, err0 := f.ReadString(module, addressPtr, addressLen)
	if err0 == 0 || len(address) != 0 {
		contractAddress = ethCommon.HexToAddress(address)
	}

	contract := bind.NewBoundContract(contractAddress, parsedAbi, client, client, client)
	_, err0 = f.handleBoundContractSize(module, contractPtr, methodsSizePtr, client, parsedAbi, contract)

	return err0
}

func (f *Factory) W_ethNewContract(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	methodsPtr uint32,
) errno.Error {
	client, err := f.getClient(clientId)
	if err != 0 {
		return err
	}

	contract, err := client.getContract(contractId)
	if err != 0 {
		return err
	}

	var methodList []string
	for method := range contract.methods {
		methodList = append(methodList, method)
	}

	return f.WriteStringSlice(module, methodsPtr, methodList)
}

func (f *Factory) W_ethGetContractMethodSize(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	methodPtr,
	methodSize,
	inputSizePtr,
	outputSizePtr uint32,
) errno.Error {
	client, err := f.getClient(clientId)
	if err != 0 {
		return err
	}

	contract, err := client.getContract(contractId)
	if err != 0 {
		return err
	}

	method, err := f.ReadString(module, methodPtr, methodSize)
	if err != 0 {
		return err
	}

	contract.methodsLock.RLock()
	contractMethod, ok := contract.methods[method]
	contract.methodsLock.RUnlock()
	if !ok {
		return errno.ErrorEthereumContractMethodNotFound
	}

	if err := f.WriteStringSliceSize(module, inputSizePtr, contractMethod.inputs); err != 0 {
		return err
	}

	return f.WriteStringSliceSize(module, outputSizePtr, contractMethod.outputs)
}

func (f *Factory) W_ethGetContractMethod(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	methodPtr,
	methodSize,
	inputPtr,
	outputPtr uint32,
) errno.Error {
	client, err := f.getClient(clientId)
	if err != 0 {
		return err
	}

	contract, err := client.getContract(contractId)
	if err != 0 {
		return err
	}

	method, err := f.ReadString(module, methodPtr, methodSize)
	if err != 0 {
		return err
	}

	contract.methodsLock.RLock()
	contractMethod, ok := contract.methods[method]
	contract.methodsLock.RUnlock()
	if !ok {
		return errno.ErrorEthereumContractMethodNotFound
	}

	err = f.WriteStringSlice(module, inputPtr, contractMethod.inputs)
	if err != 0 {
		return err
	}

	return f.WriteStringSlice(module, outputPtr, contractMethod.outputs)
}

func (f *Factory) W_ethCallContractSize(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	methodPtr,
	methodSize,
	inputsPtr,
	inputsSize,
	outPutSizePtr uint32,
) errno.Error {
	client, err0 := f.getClient(clientId)
	if err0 != 0 {
		return err0
	}

	contract, err0 := client.getContract(contractId)
	if err0 != 0 {
		return err0
	}

	methodString, err0 := f.ReadString(module, methodPtr, methodSize)
	if err0 != 0 {
		return err0
	}

	method, ok := contract.methods[methodString]
	if !ok {
		return errno.ErrorEthereumContractMethodNotFound
	}

	if !method.constant {
		return errno.ErrorEthereumCannotCallPaidMutatorTransaction
	}

	inputsBytes, err0 := f.ReadBytesSlice(module, inputsPtr, inputsSize)
	if err0 != 0 {
		return err0
	}

	inputs, err0 := verifyInputs(inputsBytes, method)
	if err0 != 0 {
		return err0
	}

	results := make([]interface{}, 0)
	err := contract.Call(nil, &results, methodString, inputs...)
	if err != nil {
		return errno.ErrorEthereumCallContractFailed
	}

	if len(results) != len(method.outputs) {
		return errno.ErrorEthereumInvalidContractMethodOutput
	}

	var outputs [][]byte
	for idx, output := range results {
		outputType := method.outputs[idx]
		if outputType == "common.Address" {
			outputs = append(outputs, output.(ethCommon.Address).Bytes())
			continue
		}

		encoder, err := codec.Converter(outputType).Encoder()
		if err != nil {
			return errno.ErrorEthereumUnsupportedDataType
		}

		value, err := encoder(output)
		if err != nil {
			return errno.ErrorEthereumParseOutputTypeFailed
		}

		if len(value) == 0 {
			value = []byte{0}
		}

		outputs = append(outputs, value)
	}

	method.data = outputs

	return f.WriteBytesSliceSize(module, outPutSizePtr, outputs)
}

func (f *Factory) W_ethCallContract(
	ctx context.Context,
	module common.Module,
	clientId,
	contractId,
	methodPtr,
	methodSize,
	outputPtr uint32,
) errno.Error {
	client, err0 := f.getClient(clientId)
	if err0 != 0 {
		return err0
	}

	contract, err0 := client.getContract(contractId)
	if err0 != 0 {
		return err0
	}

	methodString, err0 := f.ReadString(module, methodPtr, methodSize)
	if err0 != 0 {
		return err0
	}

	method, ok := contract.methods[methodString]
	if !ok {
		return errno.ErrorEthereumContractMethodNotFound
	}

	return f.WriteBytesSlice(module, outputPtr, method.data)
}
