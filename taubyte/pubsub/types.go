package pubsub

import (
	"context"

	pubsubIface "github.com/taubyte/go-interfaces/services/substrate/pubsub"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	pubsubNode pubsubIface.Service
	parent     vm.Instance
	ctx        context.Context
}

var _ vm.Factory = &Factory{}
