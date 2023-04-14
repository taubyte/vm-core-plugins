package globals

import (
	"context"

	dbIface "github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, service dbIface.Service, helper helpers.Methods) *Factory {
	return &Factory{
		parent:       i,
		ctx:          i.Context().Context(),
		databaseNode: service,
		Methods:      helper,
	}
}

func (f *Factory) Name() string {
	return "globals"
}

func (f *Factory) Close() error {
	if f.databaseInstance != nil {
		f.databaseInstance.KV().Close()
	}
	return nil
}

func (f *Factory) Context() context.Context {
	return f.ctx
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
