package function

import (
	funcIface "github.com/taubyte/go-interfaces/services/substrate/components/http"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *FunctionHttp {
	return &FunctionHttp{
		Factory: f,
		callers: make(map[uint32]funcIface.Function),
	}
}
