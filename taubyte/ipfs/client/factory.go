package client

import (
	"github.com/taubyte/go-interfaces/services/substrate/components/ipfs"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, ipfs ipfs.Service, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), ipfsNode: ipfs, Methods: helper}
}

func (f *Factory) Name() string {
	return "ipfs"
}

func (f *Factory) Close() error {
	f.clients = nil
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	f.clients = map[uint32]*Client{}
	return nil
}
