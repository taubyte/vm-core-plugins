package ethereum

import (
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), Methods: helper, clients: make(map[uint32]*Client)}
}

func (f *Factory) Name() string {
	return "ethereum"
}

func (f *Factory) Close() error {
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
