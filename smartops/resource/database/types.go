package database

import (
	"sync"

	"bitbucket.org/taubyte/go-node-database/database"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Database struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*database.Database
}
