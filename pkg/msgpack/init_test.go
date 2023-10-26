package msgpack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"

	mypkg "msgpack-json-converter/pkg/msgpack"
)

func Test_Marshal(t *testing.T) {
	type Item struct {
		Foo string
		Nil *string
	}

	t.Run("string & nil", func(t *testing.T) {
		expect := &Item{Foo: "bar", Nil: nil}

		//correctResult, _ := msgpack.Marshal(target)

		t.Log("start")

		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		actual := &Item{}
		err = msgpack.Unmarshal(toBeTest, actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		assert.Equal(t, expect.Foo, actual.Foo)
		assert.Equal(t, expect.Nil, actual.Nil)

		t.Logf("expect %+v", expect)
		t.Logf("actual %+v", actual)
	})

}
