package client

import (
	"context"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) generateClientId() uint32 {
	f.clientsLock.Lock()
	defer func() {
		f.clientsIdToGrab += 1
		f.clientsLock.Unlock()
	}()
	return f.clientsIdToGrab
}

func (f *Factory) W_newIpfsClient(ctx context.Context, module common.Module,
	clientIdPtr uint32,
) errno.Error {
	c := &Client{
		Id:       f.generateClientId(),
		Contents: make(map[uint32]*content),
	}

	f.clientsLock.Lock()
	defer f.clientsLock.Unlock()
	f.clients[c.Id] = c

	return f.WriteLe(module, clientIdPtr, c.Id)
}
