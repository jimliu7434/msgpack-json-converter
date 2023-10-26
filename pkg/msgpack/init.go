package msgpack

import "bytes"

func Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
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
