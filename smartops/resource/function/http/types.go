package function

import (
	"sync"

	"bitbucket.org/taubyte/go-node-http/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type FunctionHttp struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*function.Function
}
