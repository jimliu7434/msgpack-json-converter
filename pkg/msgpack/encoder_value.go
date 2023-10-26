package msgpack

import (
	"errors"
	"reflect"
)

func (o *Encoder) EncodeValue(v reflect.Value) error {
	if v.IsNil() {
		return o.EncodeNil()
	}

	t := v.Type()
	switch t.Kind() {
	case reflect.String:
		return o.EncodeString(v.String())
	}
	return errors.New("not implement type")
}