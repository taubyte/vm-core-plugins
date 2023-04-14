package event

import (
	"context"
	"net/http"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_eventHttpRedirect(ctx context.Context, module common.Module, eventId uint32, urlPtr uint32, urlLen uint32, code uint32) (err errno.Error) {
	url, err := f.ReadString(module, urlPtr, urlLen)
	if err != 0 {
		return
	}

	w, err := f.getEventWriter(eventId)
	if err != 0 {
		return
	}

	r, err := f.getEventRequest(eventId)
	if err != 0 {
		return
	}

	http.Redirect(w, r, url, int(code))
	return 0
}
