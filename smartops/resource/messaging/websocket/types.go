package messaging

import (
	"sync"

	"bitbucket.org/taubyte/go-node-pubsub/messaging"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type MessagingWebSocket struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*messaging.Channel
}
