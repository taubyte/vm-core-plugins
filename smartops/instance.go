package smartOps

import (
	"errors"
	"reflect"
	"strings"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/smartops/node"
	"github.com/taubyte/vm-plugins/smartops/resource"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type pluginInstance struct {
	resourceApi
	instance  vm.Instance
	factories []vm.Factory
}

// create an instance of the plugin that  can be Loaded by a wasm instance
func (p *plugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	if Plugin() == nil {
		return nil, errors.New("initialize Plugin in first")
	}

	helperMethods := helpers.New(instance.Context().Context())

	return &pluginInstance{
		instance: instance,
		factories: []vm.Factory{
			resource.New(instance, helperMethods),
			node.New(instance, p.smartOpNode, helperMethods),
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
