package helpers

import (
	"io"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

// Writes to a buffer pointer of size
func (m *methods) Read(module common.Module,
	readMethod func(p []byte) (n int, err error),
	bufPtr, bufSize, // reader
	countPtr uint32, // reader size
) errno.Error {
	buf := make([]byte, bufSize)

	n, err0 := readMethod(buf)
	if err0 != nil && err0 != io.EOF {
		return errno.ErrorHttpReadBody
	}

	ok := module.Memory().WriteUint32Le(countPtr, uint32(n))
	if !ok {
		return errno.ErrorAddressOutOfMemory
	}

	ok = module.Memory().Write(bufPtr, buf)
	if !ok {
		return errno.ErrorAddressOutOfMemory
	}

	if err0 == io.EOF {
		return errno.ErrorEOF
	}

	return 0
}
