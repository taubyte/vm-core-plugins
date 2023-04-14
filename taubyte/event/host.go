package event

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getHttpEventHostSize(ctx context.Context, module common.Module, eventId uint32, sizePtr uint32) errno.Error {
	r, err := f.getEventRequest(eventId)
	if err != 0 {
		return err
	}

	return f.WriteStringSize(module, sizePtr, r.Host)
}

func (f *Factory) W_getHttpEventHost(ctx context.Context, module common.Module, eventId uint32, bufPtr uint32, bufSize uint32) errno.Error {
	r, err := f.getEventRequest(eventId)
	if err != 0 {
		return err
	}

	return f.WriteString(module, bufPtr, r.Host)
}
