package helpers

import (
	"github.com/ipfs/go-cid"
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

var cidSize = 64

func (m *methods) ReadCid(module common.Module, ptr uint32) (cid.Cid, errno.Error) {
	cidBytes, err0 := m.ReadBytes(module, ptr, uint32(cidSize))
	if err0 != 0 {
		return cid.Cid{}, err0
	}

	_, _cid, err := cid.CidFromBytes(cidBytes)
	if err != nil {
		return cid.Cid{}, errno.ErrorInvalidCid
	}

	return _cid, 0
}

func (m *methods) WriteCid(module common.Module, ptr uint32, value cid.Cid) errno.Error {
	// validate Cid
	_cid, err0 := cid.Parse(value)
	if err0 != nil {
		return errno.ErrorInvalidCid
	}

	return m.WriteBytes(module, ptr, _cid.Bytes())
}
