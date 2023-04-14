package client

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/convert"
)

func (f *Factory) W_setHttpRequestMethod(ctx context.Context, module common.Module,
	clientId,
	requestId,
	method uint32,
) (err errno.Error) {
	_, request, err := f.getClientAndRequest(clientId, requestId)
	if err != 0 {
		return err
	}

	var err0 error
	request.Method, err0 = convert.MethodUintToString(method)
	if err0 != nil {
		return errno.ErrorInvalidMethod
	}

	return
}

func (f *Factory) W_getHttpRequestMethod(ctx context.Context, module common.Module,
	clientId,
	requestId,
	methodPtr uint32,
) (err errno.Error) {
	_, request, err := f.getClientAndRequest(clientId, requestId)
	if err != 0 {
		return
	}

	method, err0 := convert.MethodStringToUint(request.Method)
	if err0 != nil {
		return errno.ErrorInvalidMethod
	}

	return f.WriteLe(module, methodPtr, method)
}
