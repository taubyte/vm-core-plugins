package storage

import (
	"github.com/taubyte/go-interfaces/services/substrate/components/storage"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *Storage {
	return &Storage{
		Factory: f,
		callers: make(map[uint32]storage.Storage),
	}
}
