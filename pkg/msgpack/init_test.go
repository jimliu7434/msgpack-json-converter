package msgpack_test

import (
	mypkg "msgpack-json-converter/pkg/msgpack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

func showBinary(t *testing.T, data []byte) {
	limit := 8
	t.Helper()
	for i, b := range data {
		if i > limit {
			break
		}
		t.Logf("0x%02x ", b)
	}
}

func execTestingSimpleType[T any](t *testing.T, expect T, showBin bool) {
	toBeTest, err := mypkg.Marshal(expect)
	if err != nil {
		t.Fatal(err)
	}

	if showBin {
		correctBin, _ := msgpack.Marshal(expect)
		showBinary(t, correctBin)
		t.Log("---")
		showBinary(t, []byte(toBeTest))
	}

	actual := new(T)
	err = msgpack.Unmarshal(toBeTest, actual)
	if err != nil {
		t.Fatalf("unmarshal failed: %s", err.Error())
	}

	assert.Equal(t, expect, *actual)
}
