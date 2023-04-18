package function

import (
	"sync"

	"bitbucket.org/taubyte/go-node-pubsub/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type FunctionPubSub struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*function.Function
}
