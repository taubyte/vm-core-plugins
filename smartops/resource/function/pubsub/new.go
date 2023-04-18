package function

import (
	"bitbucket.org/taubyte/go-node-pubsub/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *FunctionPubSub {
	return &FunctionPubSub{
		Factory: f,
		callers: make(map[uint32]*function.Function),
	}
}
