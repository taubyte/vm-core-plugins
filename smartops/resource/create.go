package resource

import (
	smartOpIface "github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func (f *Factory) CreateSmartOp(caller smartOpIface.SmartOpEventCaller) *common.Resource {
	r := &common.Resource{
		Id:     f.generateResourceId(),
		Caller: caller,
	}

	f.resourceLock.Lock()
	defer f.resourceLock.Unlock()
	f.resources[r.Id] = r
	return r
}
