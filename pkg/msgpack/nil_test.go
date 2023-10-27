package msgpack_test

import (
	"testing"
)

func TestMarshal_Nil(t *testing.T) {
	var expect any = nil

	// toBeTest, err := mypkg.Marshal(expect)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// var actual any
	// err = msgpack.Unmarshal(toBeTest, &actual)
	// if err != nil {
	// 	t.Fatalf("unmarshal failed: %s", err.Error())
	// }

	// assert.Equal(t, expect, actual)
	execTestingSimpleType[any](t, expect, false)
}
