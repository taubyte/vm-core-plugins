package smartOps

import (
	"errors"
	"fmt"

	smartOpIface "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Instance interface {
	resourceApi
}

type resourceApi interface {
	CreateSmartOp(caller smartOpIface.SmartOpEventCaller) *common.Resource
}

var With = func(pi vm.PluginInstance) (Instance, error) {
	_pi, ok := pi.(*pluginInstance)
	if !ok {
		return nil, fmt.Errorf("type %T is not a Taubyte plugin instance", pi)
	}

	err := _pi.LoadAPIs()
	if err != nil {
		return nil, err
	}

	return _pi, nil
}

func (i *pluginInstance) LoadAPIs() error {
	var ok bool
	for _, factory := range i.factories {
		switch factory.Name() {
		case "resource":
			i.resourceApi, ok = factory.(resourceApi)
			if !ok {
				return fmt.Errorf("factory `%s` not of type `resourceApi`", factory.Name())
			}
		}
	}

	if i.resourceApi == nil {
		return errors.New("resourceApi not discovered")
	}

	return nil
}
