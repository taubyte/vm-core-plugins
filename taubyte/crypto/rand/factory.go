package rand

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
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

func (f *Factory) Context() context.Context {
	return f.ctx
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
