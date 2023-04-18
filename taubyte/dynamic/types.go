package dynamic

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	instance vm.Instance
	ctx      context.Context

	moduleInstanceLock sync.RWMutex
	moduleInstances    map[uint32]vm.ModuleInstance
	moduleInstanceIds  atomic.Uint32

	functionInstanceLock sync.RWMutex
	functionInstances    map[uint32]vm.FunctionInstance
	functionInstanceIds  atomic.Uint32
}
