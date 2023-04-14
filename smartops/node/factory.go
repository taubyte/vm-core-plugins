package node

import (
	"context"

	smartOpIface "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, service smartOpIface.Service, helper helpers.Methods) *Factory {
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

func (f *Factory) Context() context.Context {
	return f.ctx
}
