package website

import (
	"bitbucket.org/taubyte/go-node-http/website"
	"github.com/taubyte/vm-plugins/smartops/common"
)

func New(f common.Factory) *Website {
	return &Website{
		Factory: f,
		callers: make(map[uint32]*website.Website),
	}
}
