package fifo

import (
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, helper helpers.Methods) *Factory {
	return &Factory{parent: i, ctx: i.Context().Context(), Methods: helper, fifoMap: make(map[uint32]*Fifo, 0)}
}

func (f *Factory) Name() string {
	return "fifo"
}

func (f *Factory) Close() error {
	f.fifoMap = nil
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	return nil
}
