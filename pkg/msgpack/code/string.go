package msgpackcode

import (
	"encoding/binary"
	"errors"
)

const (
	_PrefixStr   = 0b10100000
	_PrefixStr8  = 0xd9
	_PrefixStr16 = 0xda
	_PrefixStr32 = 0xdb
)

const (
	_ShortStrLen = 16         // 2^4
	_Str8Len     = 256        // 2^8
	_Str16Len    = 65536      // 2^16
	_Str32Len    = 4294967296 // 2^32
)

var StrFunc = func(length int) ([]byte, error) {
	switch {
	case length < _ShortStrLen:
		return []byte{_PrefixStr + byte(length)}, nil
	case length < _Str8Len:
		return []byte{_PrefixStr8, byte(length)}, nil
	case length < _Str16Len:
		b := []byte{_PrefixStr16}
		bs := make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(length))
		b = append(b, bs...)
		return b, nil
	case length < _Str32Len:
		b := []byte{_PrefixStr32}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(length))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too long string")
	}
}
