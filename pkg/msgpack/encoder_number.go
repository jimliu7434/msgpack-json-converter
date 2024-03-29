package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeInt(v int64) error {
	bytesWithPrifix, err := msgpackcode.GetFirstBytesInt(v)
	if err != nil {
		return err
	}

	_, err = o.w.Write(bytesWithPrifix)
	return err
}

func (o *Encoder) EncodeFloat(v float64) error {
	bytesWithPrifix, err := msgpackcode.GetFirstBytesFloat(v)
	if err != nil {
		return err
	}

	_, err = o.w.Write(bytesWithPrifix)
	return err
}
