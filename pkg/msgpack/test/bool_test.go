package msgpack_test

import (
	"testing"
)

func TestMarshal_Bool(t *testing.T) {
	var testCases = []struct {
		name   string
		skip   bool
		expect bool
	}{
		{
			name:   "true",
			skip:   false,
			expect: true,
		},
		{
			name:   "false",
			skip:   false,
			expect: false,
		},
	}

	for _, tc := range testCases {
		if tc.skip {
			t.SkipNow()
		}
		t.Run(tc.name, func(t *testing.T) {
			execTestingSimpleType[bool](t, tc.expect, false)
		})
	}
}
