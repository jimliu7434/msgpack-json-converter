package msgpackcode

import (
	"encoding/binary"
	"errors"
)

const (
	_PrefixFixArr = 0b10010000
	_PrefixArr16  = 0xdc
	_PrefixArr32  = 0xdd
)

var (
	_FixArrLen = int64(15)
	_Arr16Len  = int64(32767)      // 2^15 - 1
	_Arr32Len  = int64(2147483647) // 2^31 - 1
)

func ArrayFunc(l int64) ([]byte, error) {
	switch {
	case l <= _FixArrLen:
		return []byte{_PrefixFixArr + byte(l)}, nil
	case l <= _Arr16Len:
		b := []byte{_PrefixArr16}
		bs := make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(l))
		b = append(b, bs...)
		return b, nil
	case l <= _Arr32Len:
		b := []byte{_PrefixArr32}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(l))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too long array")
	}
}
