package globals

import (
	iface "github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) kv() (iface.KV, errno.Error) {
	if f.databaseInstance == nil {
		var err error
		f.databaseInstance, err = f.databaseNode.Global(f.parent.Context().Project())
		if err != nil {
			return nil, errno.ErrorDatabaseCreateFailed
		}
	}

	return f.databaseInstance.KV(), 0
}
