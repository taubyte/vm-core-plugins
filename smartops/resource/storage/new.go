package storage

import (
	"bitbucket.org/taubyte/go-node-storage/storage"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *Storage {
	return &Storage{
		Factory: f,
		callers: make(map[uint32]*storage.Store),
	}
}
