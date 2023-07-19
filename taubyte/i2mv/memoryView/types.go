package memoryView

import (
	"context"
	"sync"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent      vm.Instance
	ctx         context.Context
	memoryViews map[uint32]*MemoryView
	mvLock      sync.RWMutex
	idsToGrab   uint32
}

type MemoryView struct {
	id       uint32
	size     uint32
	bufPtr   uint32
	closable bool
	module   vm.Module
}
