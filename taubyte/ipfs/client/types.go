package client

import (
	"context"
	"sync"

	"github.com/ipfs/go-cid"
	"github.com/taubyte/go-interfaces/services/substrate/ipfs"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	ipfsNode        ipfs.Service
	parent          vm.Instance
	ctx             context.Context
	clients         map[uint32]*Client
	clientsLock     sync.RWMutex
	clientsIdToGrab uint32
}

var _ vm.Factory = &Factory{}

type Client struct {
	Id              uint32
	contentIdToGrab uint32
	contentLock     sync.RWMutex
	Contents        map[uint32]*content
}

type content struct {
	id   uint32
	cid  cid.Cid
	file file
}

type file interface{}
