package pubsub

import (
	pubsubIface "github.com/taubyte/go-interfaces/services/substrate/pubsub"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, pubsubNode pubsubIface.Service, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), pubsubNode: pubsubNode, Methods: helper}
}

func (f *Factory) Name() string {
	return "pubsub"
}

func (f *Factory) Close() error {
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
