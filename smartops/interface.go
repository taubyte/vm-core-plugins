package smartOps

import (
	"errors"
	"fmt"

	smartOpIface "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/smartops/common"
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
		return nil, fmt.Errorf("%v of type %T is not a Taubyte plugin instance", pi, pi)
	}

	if err := _pi.LoadAPIs(); err != nil {
		return nil, err
	}

	return _pi, nil
}

func (i *pluginInstance) LoadAPIs() error {
	if i.resourceApi == nil {
		return errors.New("resourceApi not set")
	}

	return nil
}
