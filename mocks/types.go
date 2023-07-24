package mocks

import (
	"fmt"
	"net/http"
	"reflect"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/pterm/pterm"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/p2p/streams/command"
	"github.com/taubyte/p2p/streams/command/response"
	plugins "github.com/taubyte/vm-core-plugins/taubyte"
	"github.com/taubyte/vm-core-plugins/taubyte/event"
)

func init() {
	pterm.Info.Println("Initializing sdk with fake plugins")
	plugins.With = func(pi vm.PluginInstance) (plugins.Instance, error) {
		return &mockPlugin{}, nil
	}
}

type httpEvent struct {
	W http.ResponseWriter
	R *http.Request
}

type mockPlugin struct {
	AttachedFunctions     map[string]int
	CalledHttpFunctions   []httpEvent
	CalledP2PFunctions    []command.Body
	CalledPubSubFunctions []*pubsub.Message
}

func (p *mockPlugin) CreateHttpEvent(w http.ResponseWriter, r *http.Request) *event.Event {
	p.CalledHttpFunctions = append(p.CalledHttpFunctions, httpEvent{W: w, R: r})
	return &event.Event{}
}

func (p *mockPlugin) CreatePubsubEvent(msg *pubsub.Message) *event.Event {
	p.CalledPubSubFunctions = append(p.CalledPubSubFunctions, msg)
	return &event.Event{}
}

func (p *mockPlugin) CreateP2PEvent(cmd *command.Command, response response.Response) *event.Event {
	p.CalledP2PFunctions = append(p.CalledP2PFunctions, cmd.Body)
	return &event.Event{}
}

func (p *mockPlugin) AttachEvent(*event.Event) {}

func (p *mockPlugin) CheckAttached(expected map[string]int) error {
	if !reflect.DeepEqual(p.AttachedFunctions, expected) {
		return fmt.Errorf("got: %#v\nexpected: %#v", p.AttachedFunctions, expected)
	}

	return nil
}
