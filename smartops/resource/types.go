package resource

import (
	"context"
	"sync"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/smartops/common"
	"github.com/taubyte/vm-plugins/smartops/resource/database"
	functionHttp "github.com/taubyte/vm-plugins/smartops/resource/function/http"
	functionP2P "github.com/taubyte/vm-plugins/smartops/resource/function/p2p"
	functionPubSub "github.com/taubyte/vm-plugins/smartops/resource/function/pubsub"
	messagingPubSub "github.com/taubyte/vm-plugins/smartops/resource/messaging/pubsub"
	messagingWebSocket "github.com/taubyte/vm-plugins/smartops/resource/messaging/websocket"
	"github.com/taubyte/vm-plugins/smartops/resource/service"
	"github.com/taubyte/vm-plugins/smartops/resource/storage"
	"github.com/taubyte/vm-plugins/smartops/resource/website"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	*database.Database
	*functionHttp.FunctionHttp
	*functionP2P.FunctionP2P
	*functionPubSub.FunctionPubSub
	*messagingPubSub.MessagingPubSub
	*messagingWebSocket.MessagingWebSocket
	*service.Service
	*storage.Storage
	*website.Website

	helpers.Methods
	parent vm.Instance
	ctx    context.Context

	resourceLock     sync.RWMutex
	resourceIdToGrab uint32
	resources        map[uint32]*common.Resource
}

var _ vm.Factory = &Factory{}

func (f *Factory) generateResourceId() uint32 {
	f.resourceLock.Lock()
	defer func() {
		f.resourceIdToGrab += 1
		f.resourceLock.Unlock()
	}()
	return f.resourceIdToGrab
}

func (f *Factory) Load(hm vm.HostModule) (err error) {
	f.resources = map[uint32]*common.Resource{}
	f.Database = database.New(f)
	f.FunctionHttp = functionHttp.New(f)
	f.FunctionP2P = functionP2P.New(f)
	f.FunctionPubSub = functionPubSub.New(f)
	f.MessagingPubSub = messagingPubSub.New(f)
	f.MessagingWebSocket = messagingWebSocket.New(f)
	f.Service = service.New(f)
	f.Storage = storage.New(f)
	f.Website = website.New(f)
	return nil
}
