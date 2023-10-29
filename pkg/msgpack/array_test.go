package msgpack_test

import (
	"testing"
)

func TestMarshal_Array(t *testing.T) {
	showBin := false

	t.Run("single type (bool)", func(t *testing.T) {
		expect := []bool{false, false, true}
		execTestingArray[bool](t, expect, showBin)
	})

	t.Run("single type (int)", func(t *testing.T) {
		expect := []int{0, 1, 2, 3, 4}
		execTestingArray[int](t, expect, showBin)
	})

	t.Run("single type (string)", func(t *testing.T) {
		expect := []string{"", "hello", "world"}
		execTestingArray[string](t, expect, showBin)
	})

	t.Run("multi type", func(t *testing.T) {
		expect := []any{0, "hello", "world", true}
		execTestingArray[any](t, expect, showBin)
	})

	t.Run("nested array", func(t *testing.T) {
		expect := []any{[]any{0, "hello"}, "world", true}
		execTestingArray[any](t, expect, showBin)
	})
}
