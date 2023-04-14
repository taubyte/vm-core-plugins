package globals

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_putGlobalValue(
	ctx context.Context,
	module common.Module,
	namePtr, nameSize,
	// TODO, maybe send type int here
	application, function,
	valuePtr, valueSize, valueCap uint32,
) errno.Error {

	name, err0 := f.ReadString(module, namePtr, nameSize)
	if err0 != 0 {
		return err0
	}

	value, err0 := f.ReadBytes(module, valuePtr, valueSize)
	if err0 != 0 {
		return err0
	}

	path := f.getPath(application, function, name)

	db, err0 := f.kv()
	if err0 != 0 {
		return err0
	}

	err := db.Put(ctx, path, value)
	if err != nil {
		return errno.ErrorDatabaseKeyNotFound
	}

	return 0
}
