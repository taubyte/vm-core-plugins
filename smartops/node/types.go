package node

import (
	"context"

	smartOpIface "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent vm.Instance
	ctx    context.Context

	node smartOpIface.Service
}

var _ vm.Factory = &Factory{}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
