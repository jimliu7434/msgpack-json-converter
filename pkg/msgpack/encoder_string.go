package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeString(v string) error {
	codeBytes, err := msgpackcode.GetFirstBytesString(len(v))
	if err != nil {
		return err
	}

	if _, err := o.w.Write(codeBytes); err != nil {
		return err
	}
	if _, err := o.w.Write([]byte(v)); err != nil {
		return err
	}
	return nil
}
