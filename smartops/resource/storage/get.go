package storage

import (
	"bitbucket.org/taubyte/go-node-storage/storage"
	"github.com/taubyte/go-sdk/errno"
)

func (d *Storage) GetCaller(resourceId uint32) (*storage.Store, errno.Error) {
	resource, err := d.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	d.callersLock.Lock()
	defer d.callersLock.Unlock()

	db, ok := d.callers[resourceId]
	if !ok {
		db, ok = resource.Caller.(*storage.Store)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		d.callers[resourceId] = db
	}

	return db, 0
}
