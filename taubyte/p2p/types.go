package p2p

import (
	"context"
	"sync"

	"github.com/taubyte/go-interfaces/services/substrate/p2p"
	p2pIface "github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	p2pNode          p2pIface.Service
	parent           vm.Instance
	ctx              context.Context
	commandsLock     sync.RWMutex
	commandsIdToGrab uint32
	commands         map[uint32]*Command

	streamsLock sync.RWMutex
	streams     map[string]p2p.Stream

	discoverLock     sync.RWMutex
	discoverIdToGrab uint32
	discover         map[uint32][][]byte

	listenProtocol string
}

var _ vm.Factory = &Factory{}

type Command struct {
	p2pIface.Command
	Id   uint32
	Body []byte
}
