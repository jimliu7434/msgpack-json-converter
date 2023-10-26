package msgpack

import (
	"bytes"
	"io"
	"reflect"
)

type Encoder struct {
	w   io.Writer
	buf *bytes.Buffer
}

func NewEncoder(w io.Writer, buf *bytes.Buffer) *Encoder {
	e := &Encoder{
		w:   w,
		buf: buf,
	}
	return e
}

func (o *Encoder) Encode(v any) error {
	switch v := v.(type) {
	case nil:
		return o.EncodeNil()
	case string:
		return o.EncodeString(v)
		// case []byte:
		// 	return o.EncodeBytes(v)
		// case int:
		// 	return o.EncodeInt(int64(v))
		// case int64:
		// 	return o.encodeInt64Cond(v)
		// case uint:
		// 	return o.EncodeUint(uint64(v))
		// case uint64:
		// 	return o.encodeUint64Cond(v)
		// case bool:
		// 	return o.EncodeBool(v)
		// case float32:
		// 	return o.EncodeFloat32(v)
		// case float64:
		// 	return o.EncodeFloat64(v)
		// case time.Duration:
		// 	return o.encodeInt64Cond(int64(v))
		// case time.Time:
		// 	return o.EncodeTime(v)
	}
	return o.EncodeValue(reflect.ValueOf(v))
}
