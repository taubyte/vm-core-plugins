package resource

import (
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/smartops/common"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
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
