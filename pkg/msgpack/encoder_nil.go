package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeNil() error {
	_, err := o.w.Write([]byte{msgpackcode.Nil})
	return err
}
