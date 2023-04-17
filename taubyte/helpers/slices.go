package helpers

import (
	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
	"github.com/taubyte/go-sdk/utils/codec"
)

/****************************************** STRING SLICES ****************************************/

func (m *methods) ReadStringSlice(module common.Module, ptr, len uint32) ([]string, errno.Error) {
	value, ok := module.Memory().Read(ptr, len)
	if !ok {
		return nil, errno.ErrorAddressOutOfMemory
	}

	var slice []string
	err := codec.Convert(value).To(&slice)
	if err != nil {
		return nil, errno.ErrorByteConversionFailed
	}

	return slice, 0
}

func (m *methods) WriteStringSliceSize(module common.Module, ptr uint32, value []string) errno.Error {
	var encoded []byte
	err0 := codec.Convert(value).To(&encoded)
	if err0 != nil {
		return errno.ErrorByteConversionFailed
	}

	return m.WriteBytesSize(module, ptr, encoded)
}

func (m *methods) WriteStringSlice(module common.Module, ptr uint32, value []string) errno.Error {
	var encoded []byte
	err0 := codec.Convert(value).To(&encoded)
	if err0 != nil {
		return errno.ErrorByteConversionFailed
	}

	return m.WriteBytes(module, ptr, encoded)
}

/****************************************** Uint32 SLICES ****************************************/

func (m *methods) WriteUint32SliceSize(module common.Module, ptr uint32, value []uint32) errno.Error {
	var encoded []byte
	err0 := codec.Convert(value).To(&encoded)
	if err0 != nil {
		return errno.ErrorByteConversionFailed
	}

	return m.WriteBytesSize(module, ptr, encoded)
}

func (m *methods) WriteUint32Slice(module common.Module, ptr uint32, value []uint32) errno.Error {
	var encoded []byte
	err0 := codec.Convert(value).To(&encoded)
	if err0 != nil {
		return errno.ErrorByteConversionFailed
	}

	return m.WriteBytes(module, ptr, encoded)
}

// /****************************************** Uint32 SLICES ****************************************/

// func (m *methods) WriteUint64SliceSize(module common.Module, ptr uint32, value []uint64) errno.Error {

// }
