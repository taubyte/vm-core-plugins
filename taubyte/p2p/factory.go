package p2p

import (
	p2pIface "github.com/taubyte/go-interfaces/services/substrate/components/p2p"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

func New(i vm.Instance, p2pNode p2pIface.Service, helper helpers.Methods) *Factory {
	return &Factory{
		parent:  i,
		ctx:     i.Context().Context(),
		p2pNode: p2pNode,
		Methods: helper,
	}
}

func (f *Factory) Name() string {
	return "p2p"
}

func (f *Factory) Close() error {
	f.commands = nil
	f.discover = nil
	for _, stream := range f.streams {
		stream.Close()
	}
	return nil
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	f.commands = map[uint32]*Command{}
	f.streams = map[string]p2pIface.Stream{}
	f.discover = map[uint32][][]byte{}
	return nil
}
