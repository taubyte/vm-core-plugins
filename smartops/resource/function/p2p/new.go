package function

import (
	"bitbucket.org/taubyte/go-node-p2p/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *FunctionP2P {
	return &FunctionP2P{
		Factory: f,
		callers: make(map[uint32]*function.Function),
	}
}
