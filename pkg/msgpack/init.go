package msgpack

import "bytes"

func Marshal(v any) ([]byte, error) {
	enc := new(Encoder)
	var buf bytes.Buffer
	err := enc.Encode(v)
	b := buf.Bytes()
	if err != nil {
		return nil, err
	}
	return b, err
}

func Unmarshal(data []byte, v any) error {
	return nil
}
