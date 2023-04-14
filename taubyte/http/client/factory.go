package client

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), Methods: helper}
}

func (f *Factory) Name() string {
	return "client"
}

func (f *Factory) Close() error {
	f.clients = nil
	return nil
}

func (f *Factory) Context() context.Context {
	return f.ctx
}

func (f *Factory) Load(hm vm.HostModule) error {
	f.clients = map[uint32]*Client{}
	return nil
}
