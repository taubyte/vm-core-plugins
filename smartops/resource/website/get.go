package website

import (
	"bitbucket.org/taubyte/go-node-http/website"
	"github.com/taubyte/go-sdk/errno"
)

func (d *Website) GetCaller(resourceId uint32) (*website.Website, errno.Error) {
	resource, err := d.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	d.callersLock.Lock()
	defer d.callersLock.Unlock()

	db, ok := d.callers[resourceId]
	if !ok {
		db, ok = resource.Caller.(*website.Website)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		d.callers[resourceId] = db
	}

	return db, 0
}
