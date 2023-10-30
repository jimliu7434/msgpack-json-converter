package msgpackcode

import (
	"encoding/binary"
	"errors"
	"math"
)

const (
	_PrefixInt8    = 0xd0
	_PrefixInt16   = 0xd1
	_PrefixInt32   = 0xd2
	_PrefixInt64   = 0xd3
	_PrefixFloat32 = 0xca
	_PrefixFloat64 = 0xcb
)

var (
	_IntShortRange = []int64{-16, 63}                                   // positive 7bit, negative 5bit
	_Int8Range     = []int64{-128, 127}                                 // 2^7 = 128
	_Int16Range    = []int64{-32768, 32767}                             // 2^15 = 32768
	_Int32Range    = []int64{-2147483648, 2147483647}                   // 2^31 = 2147483648
	_Int64Range    = []int64{-9223372036854775808, 9223372036854775807} // 2^63 = 9223372036854775808
	_Float32Range  = []float64{-3.40282346638528859811704183484516925440e+38, 3.40282346638528859811704183484516925440e+38}
	_Float64Range  = []float64{-1.797693134862315708145274237317043567981e+308, 1.797693134862315708145274237317043567981e+308}
)

func GetFirstBytesInt(v int64) ([]byte, error) {
	switch {
	case v >= _IntShortRange[0] && v <= _IntShortRange[1]:
		if v < 0 {
			return []byte{byte(v)}, nil
		}
		return []byte{byte(v)}, nil
	case v >= _Int8Range[0] && v <= _Int8Range[1]:
		return []byte{_PrefixInt8, byte(v)}, nil
	case v >= _Int16Range[0] && v <= _Int16Range[1]:
		b := []byte{_PrefixInt16}
		bs := make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(v))
		b = append(b, bs...)
		return b, nil
	case v >= _Int32Range[0] && v <= _Int32Range[1]:
		b := []byte{_PrefixInt32}
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(v))
		b = append(b, bs...)
		return b, nil
	case v >= _Int64Range[0] && v <= _Int64Range[1]:
		b := []byte{_PrefixInt64}
		bs := make([]byte, 8)
		binary.BigEndian.PutUint64(bs, uint64(v))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too long integer")
	}
}

func GetFirstBytesFloat(v float64) ([]byte, error) {
	switch {
	case v >= _Float32Range[0] && v <= _Float32Range[1]:
		b := []byte{_PrefixFloat32}
		// to big endian IEEE 754 binary32
		bs := make([]byte, 4)
		binary.BigEndian.PutUint32(bs, math.Float32bits(float32(v)))
		b = append(b, bs...)
		return b, nil
	case v >= _Float64Range[0] && v <= _Float64Range[1]:
		b := []byte{_PrefixFloat64}
		// to big endian IEEE 754 binary64
		bs := make([]byte, 8)
		binary.BigEndian.PutUint64(bs, math.Float64bits(v))
		b = append(b, bs...)
		return b, nil
	default:
		return []byte{}, errors.New("too long float")
	}
}
