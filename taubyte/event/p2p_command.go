package event

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getP2PEventCommand(ctx context.Context, module common.Module, eventId, dataPtr uint32) errno.Error {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return err
	}

	_command, ok := data.cmd.Body["command"]
	if ok {
		data.cmd.Command, ok = _command.(string)
		if !ok {
			return errno.ErrorP2PCommandNotFound
		}
	}

	return f.WriteString(module, dataPtr, data.cmd.Command)
}

func (f *Factory) W_getP2PEventCommandSize(ctx context.Context, module common.Module, eventId, sizePtr uint32) errno.Error {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return err
	}

	_command, ok := data.cmd.Body["command"]
	if ok {
		data.cmd.Command, ok = _command.(string)
		if !ok {
			return errno.ErrorP2PCommandNotFound
		}
	}

	return f.WriteStringSize(module, sizePtr, data.cmd.Command)
}
