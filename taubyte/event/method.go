package event

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getHttpEventMethodSize(ctx context.Context, module common.Module, eventId uint32, sizePtr uint32) errno.Error {
	r, err := f.getEventRequest(eventId)
	if err != 0 {
		return err
	}

	return f.WriteStringSize(module, sizePtr, r.Method)
}

func (f *Factory) W_getHttpEventMethod(ctx context.Context, module common.Module, eventId uint32, bufPtr uint32, bufSize uint32) errno.Error {
	e, err := f.getEvent(eventId)
	if err != 0 {
		return err
	}

	m := []byte(e.http.r.Method)
	if int(bufSize) != len(m) {
		return errno.ErrorBufferTooSmall
	}

	return f.WriteBytes(module, bufPtr, m)
}
