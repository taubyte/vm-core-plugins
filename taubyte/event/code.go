package event

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_eventHttpRetCode(ctx context.Context, module common.Module, eventId uint32, code uint32) errno.Error {
	w, err := f.getEventWriter(eventId)
	if err != 0 {
		return err
	}

	w.WriteHeader(int(code))

	return 0
}
