package database

import (
	"github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *Database {
	return &Database{
		Factory: f,
		callers: make(map[uint32]database.Database),
	}
}
