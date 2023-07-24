package function

import (
	"sync"

	"github.com/taubyte/go-interfaces/services/substrate/components/pubsub"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

type FunctionPubSub struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]pubsub.Serviceable
}
