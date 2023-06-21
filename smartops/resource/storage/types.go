package storage

import (
	"sync"

	"github.com/taubyte/go-interfaces/services/substrate/storage"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Storage struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]storage.Storage
}
