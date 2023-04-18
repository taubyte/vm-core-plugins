package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) generateClientId() uint32 {
	f.clientsLock.Lock()
	defer func() {
		f.clientsIdToGrab += 1
		f.clientsLock.Unlock()
	}()
	return f.clientsIdToGrab
}

func (f *Factory) getClient(clientId uint32) (*Client, errno.Error) {
	f.clientsLock.RLock()
	defer f.clientsLock.RUnlock()
	if client, ok := f.clients[clientId]; ok {
		return client, 0
	}

	return nil, errno.ErrorClientNotFound
}

func (f *Factory) W_ethNew(ctx context.Context, module common.Module,
	clientIdPtr,
	urlPtr,
	urlLen uint32,
) errno.Error {
	url, err0 := f.ReadString(module, urlPtr, urlLen)
	if err0 != 0 {
		return err0
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		return errno.ErrorEthereumNewClient
	}

	c := Client{
		Id:        f.generateClientId(),
		Client:    client,
		blocks:    make(map[uint64]*Block),
		contracts: make(map[uint32]*Contract),
	}

	f.clientsLock.Lock()
	defer f.clientsLock.Unlock()
	f.clients[c.Id] = &c

	return f.WriteUint32Le(module, clientIdPtr, c.Id)
}

func (f *Factory) W_ethCloseClient(
	ctx context.Context,
	module common.Module,
	clientId uint32,
) errno.Error {
	client, err := f.getClient(clientId)
	if err != 0 {
		return err
	}

	client.Close()
	delete(f.clients, clientId)

	return 0
}
