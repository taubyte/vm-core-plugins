package event

import (
	"context"
	"encoding/json"

	common "github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/go-sdk/errno"
)

func (f *Factory) W_getP2PEventData(ctx context.Context, module common.Module, eventId uint32, bufPtr uint32) errno.Error {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return err
	}

	if data == nil {
		return errno.ErrorNilAddress
	}

	return f.WriteBytes(module, bufPtr, data.marshalledData)
}

func (f *Factory) W_getP2PEventDataSize(ctx context.Context, module common.Module, eventId uint32, sizePtr uint32) errno.Error {
	data, err := f.getP2PEventData(eventId)
	if err != 0 {
		return err
	}

	_data, ok := data.cmd.Body["data"]
	if !ok {
		var err0 error
		data.marshalledData, err0 = json.Marshal(data.cmd.Body)
		if err0 != nil {
			return errno.ErrorMarshalDataFailed
		}

		return f.WriteBytesSize(module, sizePtr, data.marshalledData)
	}

	data.marshalledData, ok = _data.([]byte)
	if !ok {
		return errno.ErrorMarshalDataFailed
	}

	return f.WriteBytesSize(module, sizePtr, data.marshalledData)
}
