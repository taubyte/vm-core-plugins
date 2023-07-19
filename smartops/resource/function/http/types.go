package function

import (
	"sync"

	funcIface "github.com/taubyte/go-interfaces/services/substrate/http"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

type FunctionHttp struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]funcIface.Function
}
