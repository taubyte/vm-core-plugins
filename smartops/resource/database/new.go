package database

import (
	"github.com/taubyte/go-interfaces/services/substrate/components/database"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *Database {
	return &Database{
		Factory: f,
		callers: make(map[uint32]database.Database),
	}
}
