package service

import (
	"bitbucket.org/taubyte/go-node-p2p/service"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *Service {
	return &Service{
		Factory: f,
		callers: make(map[uint32]*service.Service),
	}
}
