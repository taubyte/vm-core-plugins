package fifo

import (
	"container/list"
	"context"
	"sync"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent    vm.Instance
	ctx       context.Context
	fifoMap   map[uint32]*Fifo
	fifoLock  sync.RWMutex
	idsToGrab uint32
}

type Fifo struct {
	id         uint32
	readCloser bool
	list       *list.List
}
