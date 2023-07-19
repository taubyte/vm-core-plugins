package website

import (
	webIface "github.com/taubyte/go-interfaces/services/substrate/http"
	"github.com/taubyte/vm-core-plugins/smartops/common"
)

func New(f common.Factory) *Website {
	return &Website{
		Factory: f,
		callers: make(map[uint32]webIface.Website),
	}
}
