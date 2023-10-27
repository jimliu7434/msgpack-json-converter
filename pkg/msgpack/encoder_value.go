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

func (o *Encoder) EncodeValue(v reflect.Value) error {
	if nilable(v.Kind()) && v.IsNil() {
		return o.EncodeNil()
	}

	t := v.Type()
	log.Printf("type: %v", t.Kind())
	switch t.Kind() {
	case reflect.String:
		return o.EncodeString(v.String())
	}
	return errors.New("not implement type")
}
