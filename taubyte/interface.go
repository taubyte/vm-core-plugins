package taubyte

import (
	"errors"
	"fmt"
	"net/http"

	"bitbucket.org/taubyte/p2p/streams/command"
	"bitbucket.org/taubyte/p2p/streams/command/response"
	pubsub "github.com/libp2p/go-libp2p-pubsub"

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
	CreateP2PEvent(cmd *command.Command, response response.Response) *event.Event
}

var With = func(pi vm.PluginInstance) (Instance, error) {
	_pi, ok := pi.(*pluginInstance)
	if !ok {
		return nil, fmt.Errorf("type %T is not a Taubyte plugin instance", pi)
	}
	err := _pi.LoadAPIs()
	if err != nil {
		return nil, err
	}

	return _pi, nil
}

func (i *pluginInstance) LoadAPIs() error {
	var ok bool
	for _, factory := range i.factories {
		switch factory.Name() {
		case "event":
			i.eventApi, ok = factory.(eventApi)
			if !ok {
				return fmt.Errorf("factory `%s` not of type `eventApi`", factory.Name())
			}
		}
	}

	if i.eventApi == nil {
		return errors.New("eventApi not discovered")
	}

	return nil
}
