package function

import (
	"github.com/taubyte/go-interfaces/services/substrate/pubsub"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *FunctionPubSub {
	return &FunctionPubSub{
		Factory: f,
		callers: make(map[uint32]pubsub.Serviceable),
	}
}
