package resource

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/smartops/common"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

var _ common.Factory = &Factory{}

func New(i vm.Instance, helper helpers.Methods) *Factory {
	f := &Factory{
		parent:  i,
		ctx:     i.Context().Context(),
		Methods: helper,
	}

	return f
}

func (f *Factory) Name() string {
	return "resource"
}

func (f *Factory) Close() error {
	f.resources = nil
	return nil
}

func (f *Factory) Context() context.Context {
	return f.ctx
}
