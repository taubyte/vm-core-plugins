package messaging

import (
	messaging "github.com/taubyte/go-interfaces/services/substrate/components/pubsub"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *MessagingPubSub {
	return &MessagingPubSub{
		Factory: f,
		callers: make(map[uint32]messaging.Channel),
	}
}
