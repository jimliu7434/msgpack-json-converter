package msgpackcode

import (
	"encoding/binary"
	"errors"
)

const (
	Nil     = 0xc0
	False   = 0xc2
	True    = 0xc3
	Float32 = 0xca
	Float64 = 0xcb
	Uint8   = 0xcc
	Uint16  = 0xcd
	Uint32  = 0xce
	Uint64  = 0xcf
	Int8    = 0xd0
	Int16   = 0xd1
	Int32   = 0xd2
	Int64   = 0xd3
	Str     = 0b10100000
	Str8    = 0xd9
	Str16   = 0xda
	Str32   = 0xdb
	Array16 = 0xdc
	Array32 = 0xdd
	Map16   = 0xde
	Map32   = 0xdf
)

const (
	_ShortStrLen = 2 ^ 4  // 16
	_StrLen      = 2 ^ 8  // 256
	_Str16Len    = 2 ^ 16 // 65536
	_Str32Len    = 2 ^ 32 // 4294967296
)

var StrFunc = func(length int) ([]byte, error) {
	switch {
	case length < _ShortStrLen:
		return []byte{Str + byte(length)}, nil
	case length < _StrLen:
		return []byte{Str8, byte(length)}, nil
	case length < _Str16Len:
		b := []byte{Str16}
		bs := make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(length))
		b = append(b, bs...)
		return b, nil
	case length < _Str32Len:
		b := []byte{Str32}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(length))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too long string")
	}
}
