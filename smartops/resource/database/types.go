package database

import (
	"sync"

	"github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Database struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]database.Database
}
