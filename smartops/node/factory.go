package node

import (
	"github.com/taubyte/go-interfaces/services/substrate"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, service substrate.Service, helper helpers.Methods) *Factory {
	f := &Factory{
		parent:  i,
		ctx:     i.Context().Context(),
		Methods: helper,
		node:    service,
	}

	return f
}

func (f *Factory) Name() string {
	return "node"
}

func (f *Factory) Close() error {
	return nil
}
