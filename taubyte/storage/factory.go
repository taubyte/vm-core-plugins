package storage

import (
	"context"

	storageIface "github.com/taubyte/go-interfaces/services/substrate/storage"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, storageNode storageIface.Service, helper helpers.Methods) *Factory {
	return &Factory{
		parent:      i,
		ctx:         i.Context().Context(),
		storageNode: storageNode,
		storages:    make(map[uint32]*Storage),
		version:     make(map[string]string),
		contents:    make(map[uint32]*content),
		Methods:     helper,
	}
}

func (f *Factory) Name() string {
	return "storage"
}

func (f *Factory) Close() error {
	f.storagesLock.Lock()
	defer f.storagesLock.Unlock()
	for _, storage := range f.storages {
		storage.Close()
	}

	return nil
}

func (f *Factory) Context() context.Context {
	return f.ctx
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
