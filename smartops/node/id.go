package node

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getNodeId(ctx context.Context, module vm.Module, cidPtr uint32) errno.Error {
	return f.WriteCid(module, cidPtr, peer.ToCid(f.node.Node().ID()))
}
