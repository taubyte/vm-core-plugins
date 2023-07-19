package service

import (
	"sync"

	service "github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

type Service struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]service.ServiceResource
}
