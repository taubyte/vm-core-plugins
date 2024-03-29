package rand

import (
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), Methods: helper}
}

func (f *Factory) Name() string {
	return "crypto_rand"
}

func (f *Factory) Close() error {
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
