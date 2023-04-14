package event

import (
	"context"

	"bitbucket.org/taubyte/p2p/streams/command"
	"bitbucket.org/taubyte/p2p/streams/command/response"
	common "github.com/taubyte/go-interfaces/vm"
	sdkCommon "github.com/taubyte/go-sdk/common"
	"github.com/taubyte/go-sdk/errno"
)

type P2PData struct {
	cmd            *command.Command
	marshalledData []byte
	protocol       string
	response       response.Response
}

func (f *Factory) CreateP2PEvent(cmd *command.Command, response response.Response) *Event {
	e := &Event{
		Id:   f.generateEventId(),
		Type: sdkCommon.EventTypeP2P,
		p2p: &P2PData{
			cmd:      cmd,
			response: response,
		},
	}

	f.eventsLock.Lock()
	defer f.eventsLock.Unlock()
	f.events[e.Id] = e
	return e
}

func (f *Factory) getP2PEventData(eventId uint32) (*P2PData, errno.Error) {
	e, err := f.getEvent(eventId)
	if err != 0 {
		return nil, err
	}

	if e.p2p == nil {
		return nil, errno.ErrorNilAddress
	}

	return e.p2p, 0
}

func (f *Factory) W_writeP2PResponse(ctx context.Context, module common.Module, eventId, bufPtr, bufSize uint32) (err errno.Error) {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return
	}

	data.response["data"], err = f.ReadBytes(module, bufPtr, bufSize)

	return
}
