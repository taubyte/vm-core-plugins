package service

import (
	"sync"

	"bitbucket.org/taubyte/go-node-p2p/service"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Service struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*service.Service
}
