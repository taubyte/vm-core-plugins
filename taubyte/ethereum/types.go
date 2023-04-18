package ethereum

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent          vm.Instance
	ctx             context.Context
	clients         map[uint32]*Client
	clientsLock     sync.RWMutex
	clientsIdToGrab uint32
}

var _ vm.Factory = &Factory{}

type Client struct {
	*ethclient.Client
	Id                uint32
	blocks            map[uint64]*Block
	blocksLock        sync.RWMutex
	contracts         map[uint32]*Contract
	contractsLock     sync.RWMutex
	contractsIdToGrab uint32
}

type Block struct {
	*types.Block
	transactions     map[uint32]*Transaction
	transactionsLock sync.RWMutex
	Id               uint64
}

type Transaction struct {
	*types.Transaction
	Id uint32
}

type Contract struct {
	*bind.BoundContract
	clientId           uint32
	Id                 uint32
	methods            map[string]*contractMethod
	methodsLock        sync.RWMutex
	transactions       map[uint32]*Transaction
	transactionsLock   sync.RWMutex
	transactionsToGrab uint32
}

type contractMethod struct {
	inputs   []string
	outputs  []string
	constant bool
	data     [][]byte
}
