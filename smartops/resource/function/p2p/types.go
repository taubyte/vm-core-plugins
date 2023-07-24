package function

import (
	"sync"

	"github.com/taubyte/go-interfaces/services/substrate/components/p2p"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

type FunctionP2P struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]p2p.Serviceable
}
