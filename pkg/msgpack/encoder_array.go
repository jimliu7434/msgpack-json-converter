package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
	"reflect"
)

func (o *Encoder) EncodeArray(v reflect.Value) error {
	firstBytes, err := msgpackcode.GetFirstBytesArray(int64(v.Len()))
	if err != nil {
		return err
	}
	o.w.Write(firstBytes)
	for i := 0; i < v.Len(); i++ {
		err = o.EncodeValue(v.Index(i))
		if err != nil {
			return err
		}
	}
	return nil
}
