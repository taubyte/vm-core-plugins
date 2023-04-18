package storage

import (
	"sync"

	"bitbucket.org/taubyte/go-node-storage/storage"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Storage struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*storage.Store
}
