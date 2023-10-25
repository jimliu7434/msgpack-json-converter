package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeString(v string) error {
	bytes := []byte{}
	bytes = append(bytes, byte(msgpackcode.Str32))
	bytes = append(bytes, ([]byte(v))...)

	_, err := o.w.Write(bytes)
	return err
}
