package msgpack

import (
	"errors"
	"log"
	"reflect"
)

func nilable(kind reflect.Kind) bool {
	switch kind {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func (o *Encoder) EncodeValue(val reflect.Value) error {
	if val.Kind() == 0x00 {
		// when val is "nil" & val's type is "any"
		return o.EncodeNil()
	} else if nilable(val.Kind()) && val.IsNil() {
		return o.EncodeNil()
	}

	v := val
	t := v.Type()
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		v = v.Elem()
		t = v.Type()
	}

	if nilable(v.Kind()) && v.IsNil() {
		return o.EncodeNil()
	}

	log.Printf("type: %v", t.Kind())

	switch t.Kind() {
	case reflect.String:
		return o.EncodeString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return o.EncodeInt(v.Int())
	case reflect.Bool:
		return o.EncodeBool(v.Bool())
	case reflect.Float32, reflect.Float64:
		return o.EncodeFloat(v.Float())
	case reflect.Array, reflect.Slice:
		return o.EncodeArray(v)
	// case reflect.Map:
	// 	return o.EncodeMap(v)
	case reflect.Struct:
		return o.EncodeValue(v)
	}
	return errors.New("not implement type")
}
