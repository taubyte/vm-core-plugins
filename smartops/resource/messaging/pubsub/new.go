package messaging

import (
	"bitbucket.org/taubyte/go-node-pubsub/messaging"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *MessagingPubSub {
	return &MessagingPubSub{
		Factory: f,
		callers: make(map[uint32]*messaging.Channel),
	}
}
