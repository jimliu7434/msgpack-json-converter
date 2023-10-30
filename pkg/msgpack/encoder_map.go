package msgpack

import (
	msgpackcode "msgpack-json-converter/pkg/msgpack/code"
	"reflect"
)

func (o *Encoder) EncodeMap(v reflect.Value) error {
	ln := v.Len()
	firstBytes, err := msgpackcode.GetFirstBytesMap(int64(ln))
	if err != nil {
		return err
	}
	o.w.Write(firstBytes)
	for _, k := range v.MapKeys() {
		propName := k.String()
		propVal := v.MapIndex(k)
		o.writePropKV(propName, propVal)
	}
	return nil
}

func (o *Encoder) EncodeStruct(v reflect.Value) error {
	ln := v.NumField()
	firstBytes, err := msgpackcode.GetFirstBytesMap(int64(ln))
	if err != nil {
		return err
	}
	o.w.Write(firstBytes)
	for i := 0; i < ln; i++ {
		propName := v.Type().Field(i).Name
		propVal := v.Field(i)
		o.writePropKV(propName, propVal)
	}
	return nil
}

func (o *Encoder) writePropKV(propName string, propVal reflect.Value) error {
	// write key
	if err := o.EncodeString(propName); err != nil {
		return err
	}

	// write val
	if nilable(propVal.Kind()) && propVal.IsNil() {
		if err := o.EncodeNil(); err != nil {
			return err
		}
	}
	if propVal.Kind() == reflect.Ptr {
		propVal = propVal.Elem()
	}
	if nilable(propVal.Kind()) && propVal.IsNil() {
		if err := o.EncodeNil(); err != nil {
			return err
		}
	}
	if err := o.EncodeValue(propVal); err != nil {
		return err
	}
	return nil
}
