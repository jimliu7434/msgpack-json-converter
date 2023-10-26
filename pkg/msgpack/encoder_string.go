package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
)

func (o *Encoder) EncodeString(v string) error {
	codeBytes, err := msgpackcode.StrFunc(len(v))
	if err != nil {
		return err
	}

	bytes := []byte{}
	bytes = append(bytes, codeBytes...)
	bytes = append(bytes, ([]byte(v))...)

	_, err = o.w.Write(bytes)
	return err
}
