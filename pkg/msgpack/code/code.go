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

var StrFunc = func(length int) ([]byte, error) {

	switch {
	case length < 32:
		return []byte{Str + byte(length)}, nil
	case length < 256:
		return []byte{Str8, byte(length)}, nil
	case length < 65536:
		bs := []byte{Str16}
		binary.BigEndian.PutUint16(bs, uint16(length))
		return bs, nil
	case length < 4294967296:
		bs := []byte{Str32}
		binary.BigEndian.PutUint32(bs, uint32(length))
		return bs, nil
	default:
		return []byte{}, errors.New("too long string")
	}
}
