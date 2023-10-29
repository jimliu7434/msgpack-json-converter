package msgpack_test

import (
	"testing"
)

func TestMarshal_Nil(t *testing.T) {
	showBin := false

	t.Run("nil ptr", func(t *testing.T) {
		var expect *string = nil
		execTestingSimpleType[*string](t, expect, showBin)
	})

	t.Run("nil interface", func(t *testing.T) {
		var expect any = nil
		execTestingSimpleType[any](t, expect, showBin)
	})
}
