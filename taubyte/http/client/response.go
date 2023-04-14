package client

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_readHttpResponseBody(ctx context.Context, module common.Module,
	clientId, requestId,
	bufPtr, bufSize,
	countPtr uint32,
) (err errno.Error) {
	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return
	}

	_reader := response.Body.Read
	return f.Read(module, _reader, bufPtr, bufSize, countPtr)
}

func (f *Factory) W_closeHttpResponseBody(ctx context.Context, module common.Module,
	clientId, requestId uint32,
) (err errno.Error) {
	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return
	}

	err0 := response.Body.Close()
	if err0 != nil {
		return errno.ErrorCloseBody
	}

	return
}
