package client

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getHttpResponseHeaderSize(ctx context.Context, module common.Module,
	clientId,
	requestId,
	keyPtr, keyLen,
	sizePtr uint32,
) errno.Error {
	key, err := f.ReadString(module, keyPtr, keyLen)
	if err != 0 {
		return err
	}

	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return err
	}

	return f.WriteStringSliceSize(module, sizePtr, response.Header.Values(key))
}

func (f *Factory) W_getHttpResponseHeader(ctx context.Context, module common.Module,
	clientId,
	requestId,
	keyPtr, keyLen,
	headerPtr uint32,
) (err errno.Error) {
	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return
	}

	key, err := f.ReadString(module, keyPtr, keyLen)
	if err != 0 {
		return
	}

	return f.WriteStringSlice(module, headerPtr, response.Header.Values(key))
}

func (f *Factory) W_getHttpResponseHeaderKeysSize(ctx context.Context, module common.Module,
	clientId,
	requestId,
	sizePtr uint32,
) errno.Error {
	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return err
	}

	var keys = make([]string, 0)
	for k := range response.Header {
		keys = append(keys, k)
	}

	return f.WriteStringSliceSize(module, sizePtr, keys)
}

func (f *Factory) W_getHttpResponseHeaderKeys(ctx context.Context, module common.Module,
	clientId,
	requestId,
	keysPtr, keysSize uint32,
) errno.Error {
	response, err := f.getResponse(clientId, requestId)
	if err != 0 {
		return err
	}

	var keys = make([]string, 0)
	for k := range response.Header {
		keys = append(keys, k)
	}

	return f.WriteStringSlice(module, keysPtr, keys)
}
