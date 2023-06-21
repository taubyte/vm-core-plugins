package website

import (
	"sync"

	webIface "github.com/taubyte/go-interfaces/services/substrate/http"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Website struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]webIface.Website
}
