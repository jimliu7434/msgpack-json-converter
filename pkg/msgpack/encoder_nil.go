package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeNil() error {
	_, err := o.w.Write([]byte{msgpackcode.Nil})
	return err
}

func (o *Encoder) EncodeBool(b bool) (err error) {
	if b {
		_, err = o.w.Write([]byte{msgpackcode.True})
	} else {
		_, err = o.w.Write([]byte{msgpackcode.False})
	}

	return err
}
