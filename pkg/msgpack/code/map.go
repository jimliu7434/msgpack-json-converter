package msgpackcode

import (
	"encoding/binary"
	"errors"
)

const (
	_PrefixFixMap = 0b10000000
	_PrefixMap16  = 0xde
	_PrefixMap32  = 0xdf
)

var (
	_FixMapLen = int64(15)
	_Map16Len  = int64(32767)      // 2^15 - 1
	_Map32Len  = int64(2147483647) // 2^31 - 1
)

func GetFirstBytesMap(l int64) ([]byte, error) {
	switch {
	case l <= _FixMapLen:
		return []byte{_PrefixFixMap + byte(l)}, nil
	case l <= _Map16Len:
		b := []byte{_PrefixMap16}
		bs := make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(l))
		b = append(b, bs...)
		return b, nil
	case l <= _Map32Len:
		b := []byte{_PrefixMap32}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(l))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too big map")
	}
}
