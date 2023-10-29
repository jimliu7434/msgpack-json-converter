package msgpack

import (
	"bytes"
	"io"
	"reflect"
)

type Encoder struct {
	w io.Writer
}

func NewEncoder(buf *bytes.Buffer) *Encoder {
	w := io.Writer(buf)
	e := &Encoder{
		w: w,
	}
	return e
}

func (o *Encoder) Encode(v any) error {
	switch v := v.(type) {
	case string:
		return o.EncodeString(v)
	case int:
		return o.EncodeInt(int64(v))
	case int8:
		return o.EncodeInt(int64(v))
	case int16:
		return o.EncodeInt(int64(v))
	case int32:
		return o.EncodeInt(int64(v))
	case int64:
		return o.EncodeInt(v)
	case bool:
		return o.EncodeBool(v)
	case float32:
		return o.EncodeFloat(float64(v))
	case float64:
		return o.EncodeFloat(v)
	}
	return o.EncodeValue(reflect.ValueOf(v))
}
