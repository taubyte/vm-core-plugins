package helpers

import (
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (m *methods) ReadString(module common.Module, ptr, len uint32) (string, errno.Error) {
	value, ok := module.Memory().Read(ptr, len)
	if !ok {
		return "", errno.ErrorAddressOutOfMemory
	}

	return string(value), 0
}

func (m *methods) WriteStringSize(module common.Module, ptr uint32, data string) errno.Error {
	return m.WriteUint32Le(module, ptr, uint32(len(data)))
}

func (m *methods) WriteString(module common.Module, ptr uint32, value string,
) errno.Error {
	ok := module.Memory().Write(ptr, []byte(value))
	if !ok {
		return errno.ErrorAddressOutOfMemory
	}

	return 0
}
