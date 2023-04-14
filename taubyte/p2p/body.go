package p2p

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_readCommandResponse(ctx context.Context, module common.Module,
	commandId,
	dataBuf, dataSize uint32,
) (err errno.Error) {
	cmd, err := f.getCommand(commandId)
	if err != 0 {
		return
	}

	return f.SafeWriteBytes(module, dataBuf, dataSize, cmd.Body)
}
