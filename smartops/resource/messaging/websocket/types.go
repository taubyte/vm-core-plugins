package messaging

import (
	"sync"

	messaging "github.com/taubyte/go-interfaces/services/substrate/components/pubsub"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

type MessagingWebSocket struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]messaging.Messaging
}
