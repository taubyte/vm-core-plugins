package service

import (
	service "github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Service) GetCaller(resourceId uint32) (service.ServiceResource, errno.Error) {
	resource, err := f.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	f.callersLock.Lock()
	defer f.callersLock.Unlock()

	message, ok := f.callers[resourceId]
	if !ok {
		message, ok = resource.Caller.(service.ServiceResource)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		f.callers[resourceId] = message
	}

	return message, 0
}
