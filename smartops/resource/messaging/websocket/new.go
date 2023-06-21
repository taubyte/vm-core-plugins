package messaging

import (
	messaging "github.com/taubyte/go-interfaces/services/substrate/pubsub"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *MessagingWebSocket {
	return &MessagingWebSocket{
		Factory: f,
		callers: make(map[uint32]messaging.Messaging),
	}
}
