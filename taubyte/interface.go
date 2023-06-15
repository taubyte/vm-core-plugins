package taubyte

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/taubyte/go-interfaces/p2p/streams"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/event"
)

type Instance interface {
	eventApi
}

type eventApi interface {
	AttachEvent(*event.Event)

	CreateHttpEvent(w http.ResponseWriter, r *http.Request) *event.Event
	CreatePubsubEvent(msg *pubsub.Message) *event.Event
	CreateP2PEvent(cmd streams.Command, response streams.Response) *event.Event
}

var With = func(pi vm.PluginInstance) (Instance, error) {
	_pi, ok := pi.(*pluginInstance)
	if !ok {
		debug.PrintStack()
		return nil, fmt.Errorf("%v of type %T is not a Taubyte plugin instance", pi, pi)
	}

	if err := _pi.LoadAPIs(); err != nil {
		return nil, err
	}

	return _pi, nil
}

var _ eventApi = &event.Factory{}

func (i *pluginInstance) LoadAPIs() (err error) {
	if i.eventApi == nil {
		err = errors.New("eventApi not set ")
	}

	return
}
