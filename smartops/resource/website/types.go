package website

import (
	"sync"

	"bitbucket.org/taubyte/go-node-http/website"
	"github.com/taubyte/vm-plugins/smartops/common"
)

type Website struct {
	common.Factory

	callersLock sync.RWMutex
	callers     map[uint32]*website.Website
}
