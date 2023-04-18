package function

import (
	"sync"

	"bitbucket.org/taubyte/go-node-p2p/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type FunctionP2P struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*function.Function
}
