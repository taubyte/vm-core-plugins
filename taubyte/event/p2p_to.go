package event

import (
	"context"

	"github.com/libp2p/go-libp2p/core/peer"
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getP2PEventTo(ctx context.Context, module common.Module, eventId, cidPtr uint32) errno.Error {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return err
	}

	conn, err0 := data.cmd.Connection()
	if err0 != nil {
		return err
	}

	_to := conn.LocalPeer()
	if len(_to) == 0 {
		return errno.ErrorP2PToNotFound
	}

	return f.WriteCid(module, cidPtr, peer.ToCid(_to))
}
