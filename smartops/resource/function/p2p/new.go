package function

import (
	"github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *FunctionP2P {
	return &FunctionP2P{
		Factory: f,
		callers: make(map[uint32]p2p.Serviceable),
	}
}
