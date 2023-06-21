package service

import (
	service "github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *Service {
	return &Service{
		Factory: f,
		callers: make(map[uint32]service.ServiceResource),
	}
}
