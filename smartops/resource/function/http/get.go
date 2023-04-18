package function

import (
	"bitbucket.org/taubyte/go-node-http/function"
	"github.com/taubyte/go-sdk/errno"
)

func (f *FunctionHttp) GetCaller(resourceId uint32) (*function.Function, errno.Error) {
	resource, err := f.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	f.callersLock.Lock()
	defer f.callersLock.Unlock()

	_func, ok := f.callers[resourceId]
	if !ok {
		_func, ok = resource.Caller.(*function.Function)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		f.callers[resourceId] = _func
	}

	return _func, 0
}
