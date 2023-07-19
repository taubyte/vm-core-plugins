package client

import (
	dbIface "github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, service dbIface.Service, helper helpers.Methods) *Factory {
	return &Factory{
		parent:       i,
		ctx:          i.Context().Context(),
		databaseNode: service,
		database:     make(map[uint32]*Database),
		Methods:      helper,
	}
}

func (f *Factory) Name() string {
	return "database"
}

func (f *Factory) Close() error {
	f.databaseLock.Lock()
	defer f.databaseLock.Unlock()
	for _, database := range f.database {
		database.Close()
	}

	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
