package main

import (
	"encoding/json"
	mymsgpack "msgpack-json-converter/pkg/msgpack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

type myObject struct {
	MyStr       string    `json:"my_str"`
	MyInt       int16     `json:"my_int"`
	MyBool      bool      `json:"my_bool"`
	MyFloat     float32   `json:"my_float"`
	MySubObject *myObject `json:"my_sub_object"`
}

func TestMyMsgPackMarshal(t *testing.T) {
	myJSONstr := `{"my_str":"hello","my_int":1000,"my_bool":true,"my_float":1.23,"my_sub_object":{"my_str":"world","my_int":2000,"my_bool":false,"my_float":11.2233,"my_sub_object":null}}`

	// parse to myobject
	var myObj myObject
	err := json.Unmarshal([]byte(myJSONstr), &myObj)
	if err != nil {
		// json format malformed
		t.Fatalf("json.Unmarshal failed: %v", err)
	}

	// convert myobject to msgpack
	msgpackBytes, err := mymsgpack.Marshal(&myObj)
	if err != nil {
		t.Fatalf("mymsgpack.Marshal failed: %v", err)
	}

	// use 3rd-party msgpack to verify
	correctBin, _ := msgpack.Marshal(myObj)
	assert.Equal(t, correctBin, msgpackBytes)
}
