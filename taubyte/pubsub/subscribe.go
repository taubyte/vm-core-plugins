package pubsub

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_setSubscriptionChannel(ctx context.Context, module common.Module,
	channelPtr, channelLen uint32,
) (err errno.Error) {
	channel, err := f.ReadString(module, channelPtr, channelLen)
	if err != 0 {
		return
	}

	_ctx := f.parent.Context()

	err0 := f.pubsubNode.Subscribe(_ctx.Project(), _ctx.Application(), channel)
	if err0 != nil {
		return errno.ErrorSubscribeFailed
	}

	return 0
}
