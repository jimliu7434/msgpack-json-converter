package msgpack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"

	mypkg "msgpack-json-converter/pkg/msgpack"
)

func Test_Marshal(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var expect any = nil

		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		var actual any
		err = msgpack.Unmarshal(toBeTest, &actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		assert.Equal(t, expect, actual)
	})

	t.Run("string", func(t *testing.T) {
		var testingCases = []struct {
			name        string
			skip        bool
			expect      string
			expectedErr bool
		}{
			{
				name:        "empty",
				skip:        false,
				expect:      "",
				expectedErr: false,
			},
			{
				name:        "short string",
				skip:        false,
				expect:      "foo",
				expectedErr: false,
			},
			{
				name:        "str8 string",
				skip:        false,
				expect:      "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfa",
				expectedErr: false,
			},
			{
				name:        "str16 string",
				skip:        false,
				expect:      "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfalkjsdhflkajdshfjklasdhflkajdshflkajsdhfkjasdhkjahsdkljfahsdkljfhalksdjhfalksdjhflkajsdhfkljashdfklahsdkfhasdkjfhkajdfhkajshdflkajshdflkjashdflkjahsdlkfjhasdkljfhlkadjhflkajdshflkashdfkjlashdflkajhdsfkljahsdlkjfahdslkjfhaslkdjfhalkjsdhflkajsdhflkjasdhflkjdsh",
				expectedErr: false,
			},
			{
				name:        "str32 string",
				skip:        false,
				expect:      "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfalkjsdhflkajdshfjklasdhflkajdshflkajsdhfkjasdhkjahsdkljfahsdkljfhalksdjhfalksdjhflkajsdhfkljashdfklahsdkfhasdkjfhkajdfhkajshdflkajshdflkjashdflkjahsdlkfjhasdkljfhlkadjhflkajdshflkashdfkjlashdflkajhdsfkljahsdlkjfahdslkjfhaslkdjfhalkjsdhflkajsdhflkjasdhflkjdsh",
				expectedErr: false,
			},
			{
				name:        "special character",
				skip:        false,
				expect:      "あいうえお",
				expectedErr: false,
			},
		}

		for _, tc := range testingCases {
			t.Run(tc.name, func(t *testing.T) {
				if tc.skip {
					t.SkipNow()
				}

				correctBin, _ := msgpack.Marshal(tc.expect)
				showBinary(t, correctBin)
				t.Log("---")
				expect := tc.expect
				toBeTest, err := mypkg.Marshal(expect)
				if tc.expectedErr {
					assert.Error(t, err)
					return
				}

				showBinary(t, []byte(toBeTest))

				if err != nil {
					t.Fatal(err)
				}

				var actual string
				err = msgpack.Unmarshal(toBeTest, &actual)
				if err != nil {
					t.Fatalf("unmarshal failed: %s", err.Error())
				}
				assert.Equal(t, expect, actual)
			})
		}
	})

	t.Run("number", func(t *testing.T) {
		var expect any = nil

		//correctResult, _ := msgpack.Marshal(target)

		t.Log("start")

		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		//actual := &Item{}
		var actual any
		err = msgpack.Unmarshal(toBeTest, &actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		//assert.Equal(t, expect.Foo, actual.Foo)
		//assert.Equal(t, expect.Nil, actual.Nil)

		t.Logf("expect %+v", expect)
		t.Logf("actual %+v", actual)
	})

	t.Run("object", func(t *testing.T) {
		t.SkipNow()
		//expect := &Item{Foo: "bar", Nil: nil}
		var expect any = nil

		//correctResult, _ := msgpack.Marshal(target)

		t.Log("start")

		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		//actual := &Item{}
		var actual any
		err = msgpack.Unmarshal(toBeTest, &actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		//assert.Equal(t, expect.Foo, actual.Foo)
		//assert.Equal(t, expect.Nil, actual.Nil)

		t.Logf("expect %+v", expect)
		t.Logf("actual %+v", actual)
	})

	t.Run("array", func(t *testing.T) {
		t.SkipNow()
		//expect := &Item{Foo: "bar", Nil: nil}
		var expect any = nil

		//correctResult, _ := msgpack.Marshal(target)
		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		//actual := &Item{}
		var actual any
		err = msgpack.Unmarshal(toBeTest, &actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		//assert.Equal(t, expect.Foo, actual.Foo)
		//assert.Equal(t, expect.Nil, actual.Nil)

		t.Logf("expect %+v", expect)
		t.Logf("actual %+v", actual)
	})

	t.Run("nested object", func(t *testing.T) {
		t.SkipNow()
		//expect := &Item{Foo: "bar", Nil: nil}
		var expect any = nil

		//correctResult, _ := msgpack.Marshal(target)

		toBeTest, err := mypkg.Marshal(expect)
		if err != nil {
			t.Fatal(err)
		}

		//actual := &Item{}
		var actual any
		err = msgpack.Unmarshal(toBeTest, &actual)
		if err != nil {
			t.Fatalf("unmarshal failed: %s", err.Error())
		}

		//assert.Equal(t, expect.Foo, actual.Foo)
		//assert.Equal(t, expect.Nil, actual.Nil)

		t.Logf("expect %+v", expect)
		t.Logf("actual %+v", actual)
	})
}

func showBinary(t *testing.T, data []byte) {
	limit := 16
	t.Helper()
	for i, b := range data {
		if i > limit {
			break
		}
		t.Logf("0x%02x ", b)
	}
}
