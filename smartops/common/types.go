package common

import (
	"github.com/taubyte/go-interfaces/services/substrate/smartops"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/vm-core-plugins/taubyte/event"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory interface {
	helpers.Methods
	GetEvent(resourceId uint32) (*event.Event, errno.Error)
	GetResource(resourceId uint32) (*Resource, errno.Error)
}

type Resource struct {
	Id     uint32
	Caller smartops.EventCaller
}
