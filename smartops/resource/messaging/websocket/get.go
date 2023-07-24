package messaging

import (
	messaging "github.com/taubyte/go-interfaces/services/substrate/components/pubsub"
	"github.com/taubyte/go-sdk/errno"
)

func (f *MessagingWebSocket) GetCaller(resourceId uint32) (messaging.Messaging, errno.Error) {
	resource, err := f.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	f.callersLock.Lock()
	defer f.callersLock.Unlock()

	message, ok := f.callers[resourceId]
	if !ok {
		message, ok = resource.Caller.(messaging.Messaging)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		f.callers[resourceId] = message
	}

	return message, 0
}
