package msgpack_test

import (
	mypkg "msgpack-json-converter/pkg/msgpack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
)

type Object struct {
	Int    int
	String string
	Bool   bool
	Float  float64
	Object *Object
}

func TestMarshal_Object(t *testing.T) {
	showBin := false
	t.SkipNow()

	var testCases = []struct {
		name   string
		skip   bool
		expect Object
	}{
		{
			name: "int, string, bool, float",
			skip: false,
			expect: Object{
				Int:    0,
				String: "hello",
				Bool:   true,
				Float:  1.23,
				Object: nil,
			},
		},
		{
			name: "nested object",
			skip: false,
			expect: Object{
				Int:    0,
				String: "hello",
				Bool:   true,
				Float:  1.23,
				Object: &Object{
					Int:    1,
					String: "world",
					Bool:   false,
					Float:  11.2233,
					Object: nil,
				},
			},
		},
	}

	for _, tc := range testCases {
		if tc.skip {
			t.SkipNow()
		}
		t.Run(tc.name, func(t *testing.T) {
			toBeTest, err := mypkg.Marshal(tc.expect)
			if err != nil {
				t.Fatal(err)
			}

			if showBin {
				correctBin, _ := msgpack.Marshal(tc.expect)
				showBinary(t, correctBin)
				t.Log("---")
				showBinary(t, []byte(toBeTest))
			}

			actual := &Object{}
			err = msgpack.Unmarshal(toBeTest, actual)
			if err != nil {
				t.Fatalf("unmarshal failed: %s", err.Error())
			}

			assert.Equal(t, tc.expect.Int, actual.Int)
			assert.Equal(t, tc.expect.String, actual.String)
			assert.Equal(t, tc.expect.Bool, actual.Bool)
			assert.Equal(t, tc.expect.Float, actual.Float)
			if tc.expect.Object == nil {
				assert.Nil(t, actual.Object)
			} else {
				assert.Equal(t, tc.expect.Object.Int, actual.Object.Int)
				assert.Equal(t, tc.expect.Object.String, actual.Object.String)
				assert.Equal(t, tc.expect.Object.Bool, actual.Object.Bool)
				assert.Equal(t, tc.expect.Object.Float, actual.Object.Float)
				assert.Nil(t, actual.Object.Object)
			}
		})
	}
}
