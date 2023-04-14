package taubyte

import (
	"errors"
	"reflect"
	"strings"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/dns"
	"github.com/taubyte/vm-plugins/taubyte/ethereum"
	"github.com/taubyte/vm-plugins/taubyte/globals"
	"github.com/taubyte/vm-plugins/taubyte/pubsub"
	"github.com/taubyte/vm-plugins/taubyte/self"
	"github.com/taubyte/vm-plugins/taubyte/storage"

	"github.com/taubyte/vm-plugins/taubyte/crypto/rand"
	kvdb "github.com/taubyte/vm-plugins/taubyte/database/client"
	"github.com/taubyte/vm-plugins/taubyte/event"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
	"github.com/taubyte/vm-plugins/taubyte/http/client"
	"github.com/taubyte/vm-plugins/taubyte/i2mv/fifo"
	"github.com/taubyte/vm-plugins/taubyte/i2mv/memoryView"
	ipfsClient "github.com/taubyte/vm-plugins/taubyte/ipfs/client"
	p2pClient "github.com/taubyte/vm-plugins/taubyte/p2p"
)

type pluginInstance struct {
	eventApi
	instance  vm.Instance
	factories []vm.Factory
}

// create an instance of the plugin that  can be Loaded by a wasm instance
func (p *plugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	if Plugin() == nil {
		return nil, errors.New("initialize plugin in first")
	}

	helperMethods := helpers.New(instance.Context().Context())

	return &pluginInstance{
		instance: instance,
		factories: []vm.Factory{
			event.New(instance, helperMethods),
			ethereum.New(instance, helperMethods),
			client.New(instance, helperMethods),
			ipfsClient.New(instance, p.ipfsNode, helperMethods),
			pubsub.New(instance, p.pubsubNode, helperMethods),
			storage.New(instance, p.storageNode, helperMethods),
			kvdb.New(instance, p.databaseNode, helperMethods),
			p2pClient.New(instance, p.p2pNode, helperMethods),
			dns.New(instance, helperMethods),
			self.New(instance, helperMethods),
			globals.New(instance, p.databaseNode, helperMethods),
			rand.New(instance, helperMethods),
			memoryView.New(instance, helperMethods),
			fifo.New(instance, helperMethods),
		},
	}, nil
}

func (i *pluginInstance) LoadFactory(factory vm.Factory, hm vm.HostModule) error {
	err := factory.Load(hm)
	if err != nil {
		return err
	}
	defs := make([]*vm.HostModuleFunctionDefinition, 0)
	m := reflect.ValueOf(factory)
	mT := reflect.TypeOf(factory)
	for i := 0; i < m.NumMethod(); i++ {
		mt := m.Method(i)
		mtT := mT.Method(i)
		if !strings.HasPrefix(mtT.Name, "W_") {
			continue
		}

		defs = append(defs, &vm.HostModuleFunctionDefinition{
			Name:    mtT.Name[2:],
			Handler: mt.Interface(),
		})
	}

	return hm.Functions(defs...)
}
func (i *pluginInstance) Load(hm vm.HostModule) (vm.ModuleInstance, error) {
	errs := make([]string, 0)
	for _, factory := range i.factories {
		err := i.LoadFactory(factory, hm)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	// build a composite error
	if len(errs) > 0 {
		return nil, errors.New("Load failed with:\n" + strings.Join(errs, "\n; "))
	}

	return hm.Compile()
}

func (i *pluginInstance) Close() error {
	errs := make([]string, 0)
	for _, factory := range i.factories {
		err := factory.Close()
		if err != nil {
			errs = append(errs, err.Error())
		}
	}

	// build a composite error
	if len(errs) > 0 {
		return errors.New("Close failed with:\n" + strings.Join(errs, "\n; "))
	}

	return nil
}
