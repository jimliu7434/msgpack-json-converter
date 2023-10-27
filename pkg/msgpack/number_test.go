package msgpack_test

import (
	"testing"
)

func TestMarshal_Int(t *testing.T) {
	showBin := false

	t.Run("int", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect int
		}{
			{
				name:   "zero",
				skip:   false,
				expect: 0,
			},
			{
				name:   "positive small",
				skip:   false,
				expect: 1,
			},
			{
				name:   "negative small",
				skip:   false,
				expect: -10,
			},
			{
				name:   "positive large",
				skip:   false,
				expect: 32767,
			},
			{
				name:   "negative large",
				skip:   false,
				expect: -32768,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[int](t, tc.expect, showBin)
			})
		}
	})

	t.Run("int8", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect int8
		}{
			{
				name:   "zero",
				skip:   false,
				expect: 0,
			},
			{
				name:   "positive int8",
				skip:   false,
				expect: 127,
			},
			{
				name:   "negative int8",
				skip:   false,
				expect: -128,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[int8](t, tc.expect, showBin)
			})
		}
	})

	t.Run("int16", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect int16
		}{
			{
				name:   "positive int16",
				skip:   false,
				expect: 32767,
			},
			{
				name:   "negative int16",
				skip:   false,
				expect: -32768,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[int16](t, tc.expect, showBin)
			})
		}
	})

	t.Run("int32", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect int32
		}{
			{
				name:   "positive int32",
				skip:   false,
				expect: 2147483647,
			},
			{
				name:   "negative int32",
				skip:   false,
				expect: -2147483648,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[int32](t, tc.expect, showBin)
			})
		}
	})

	t.Run("int64", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect int64
		}{
			{
				name:   "positive int64",
				skip:   false,
				expect: 9223372036854775807,
			},
			{
				name:   "negative int64",
				skip:   false,
				expect: -9223372036854775808,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[int64](t, tc.expect, showBin)
			})
		}
	})
}

func TestMarshal_Float(t *testing.T) {
	showBin := false
	t.Run("float32", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect float32
		}{
			{
				name:   "zero",
				skip:   false,
				expect: 0.0,
			},
			{
				name:   "positive float32 1",
				skip:   false,
				expect: 3.40282346638528859811704183484516925440e+38,
			},
			{
				name:   "negative float32 1",
				skip:   false,
				expect: -3.40282346638528859811704183484516925440e+38,
			},
			{
				name:   "positive float32 2",
				skip:   false,
				expect: 1.0,
			},
			{
				name:   "negative float32 2",
				skip:   false,
				expect: -1.0,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[float32](t, tc.expect, showBin)
			})
		}
	})
	t.Run("float64", func(t *testing.T) {
		var testCases = []struct {
			name   string
			skip   bool
			expect float64
		}{
			{
				name:   "positive float64",
				skip:   false,
				expect: 1.797693134862315708145274237317043567981e+308,
			},
			{
				name:   "negative float64",
				skip:   false,
				expect: -1.797693134862315708145274237317043567981e+308,
			},
		}

		for _, tc := range testCases {
			if tc.skip {
				t.SkipNow()
			}
			t.Run(tc.name, func(t *testing.T) {
				execTestingSimpleType[float64](t, tc.expect, showBin)
			})
		}
	})
}
