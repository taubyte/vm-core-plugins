package service

import (
	"bitbucket.org/taubyte/go-node-p2p/service"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Service) GetCaller(resourceId uint32) (*service.Service, errno.Error) {
	resource, err := f.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	f.callersLock.Lock()
	defer f.callersLock.Unlock()

	message, ok := f.callers[resourceId]
	if !ok {
		message, ok = resource.Caller.(*service.Service)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		f.callers[resourceId] = message
	}

	return message, 0
}
