package msgpack_test

import (
	"testing"
)

func TestMarshal_String(t *testing.T) {
	var testingCases = []struct {
		name   string
		skip   bool
		expect string
	}{
		{
			name:   "empty",
			skip:   false,
			expect: "",
		},
		{
			name:   "short string",
			skip:   false,
			expect: "foo",
		},
		{
			name:   "str8 string",
			skip:   false,
			expect: "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfa",
		},
		{
			name:   "str16 string",
			skip:   false,
			expect: "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfalkjsdhflkajdshfjklasdhflkajdshflkajsdhfkjasdhkjahsdkljfahsdkljfhalksdjhfalksdjhflkajsdhfkljashdfklahsdkfhasdkjfhkajdfhkajshdflkajshdflkjashdflkjahsdlkfjhasdkljfhlkadjhflkajdshflkashdfkjlashdflkajhdsfkljahsdlkjfahdslkjfhaslkdjfhalkjsdhflkajsdhflkjasdhflkjdsh",
		},
		{
			name:   "str32 string",
			skip:   false,
			expect: "alkjdhfklajshdlkfjahsdkjfhalskjdhfalkjsdhflkajsdhflkjahdskjfahdjklfhalksdjhfalkjsdhflkajdshfjklasdhflkajdshflkajsdhfkjasdhkjahsdkljfahsdkljfhalksdjhfalksdjhflkajsdhfkljashdfklahsdkfhasdkjfhkajdfhkajshdflkajshdflkjashdflkjahsdlkfjhasdkljfhlkadjhflkajdshflkashdfkjlashdflkajhdsfkljahsdlkjfahdslkjfhaslkdjfhalkjsdhflkajsdhflkjasdhflkjdsh",
		},
		{
			name:   "special character",
			skip:   false,
			expect: "あいうえお",
		},
	}

	for _, tc := range testingCases {
		if tc.skip {
			t.SkipNow()
		}
		t.Run(tc.name, func(t *testing.T) {
			execTestingSimpleType[string](t, tc.expect, false)
		})
	}
}
