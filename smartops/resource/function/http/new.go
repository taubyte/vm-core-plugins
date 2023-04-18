package function

import (
	"bitbucket.org/taubyte/go-node-http/function"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *FunctionHttp {
	return &FunctionHttp{
		Factory: f,
		callers: make(map[uint32]*function.Function),
	}
}
