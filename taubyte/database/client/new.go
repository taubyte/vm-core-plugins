package client

import (
	"context"

	databaseIface "github.com/taubyte/go-interfaces/services/substrate/database"
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_newDatabase(ctx context.Context,
	module common.Module,
	databaseMatchPtr, databaseMatchSize,
	idPtr uint32,
) errno.Error {

	databaseMatch, err := f.ReadString(module, databaseMatchPtr, databaseMatchSize)
	if err != 0 {
		return err
	}

	_ctx := f.parent.Context()
	databaseContext := databaseIface.Context{
		ProjectId:     _ctx.Project(),
		ApplicationId: _ctx.Application(),
		Matcher:       databaseMatch,
	}

	_database, err0 := f.databaseNode.Database(databaseContext)
	if err0 != nil {
		return errno.ErrorDatabaseCreateFailed
	}

	database := f.createDatabasePointer(_database)

	return f.WriteLe(module, idPtr, uint32(database.Id))
}
