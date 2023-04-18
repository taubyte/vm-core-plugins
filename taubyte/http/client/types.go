package client

import (
	"context"
	"net/http"
	"sync"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent          vm.Instance
	ctx             context.Context
	clientsLock     sync.RWMutex
	clientsIdToGrab uint32
	clients         map[uint32]*Client
}

var _ vm.Factory = &Factory{}

type Client struct {
	*http.Client
	Id          uint32
	reqLock     sync.RWMutex
	reqIdToGrab uint32
	reqs        map[uint32]*Request
}

type Request struct {
	*http.Request
	Id       uint32
	dataPtr  uint32
	dataSize uint32
}
