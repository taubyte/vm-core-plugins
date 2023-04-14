package helpers

import (
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (m *methods) SafeWriteBytes(module common.Module, ptr, size uint32, value []byte,
) errno.Error {
	if uint32(len(value)) != size {
		return errno.ErrorSizeMismatch
	}

	ok := module.Memory().Write(ptr, value)
	if !ok {
		return errno.ErrorAddressOutOfMemory
	}

	return 0
}
