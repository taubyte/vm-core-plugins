package client

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/vm-core-plugins/taubyte/memory"
)

func (f *Factory) W_setHttpRequestBody(ctx context.Context, module common.Module,
	clientId, requestId,
	bodyPtr, bodySize uint32,
) (err errno.Error) {
	_, req, err := f.getClientAndRequest(clientId, requestId)
	if err != 0 {
		return
	}

	req.Body = memory.New(f.ctx, module.Memory(), bodyPtr, bodySize)

	return 0
}
