package rsltutil

import (
	"errors"
	"testing"

	"github.com/SSripilaipong/go-common/rslt"
	"github.com/stretchr/testify/assert"
)

func TestResultOf2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		first := rslt.Value("cmd")
		second := rslt.Value("input")

		result := ResultOf2(first, second)

		assert.True(t, result.IsOk())
		x1, x2 := result.Value().Return()
		assert.Equal(t, "cmd", x1)
		assert.Equal(t, "input", x2)
	})

	t.Run("CommandError", func(t *testing.T) {
		err := errors.New("cmd err")
		result := ResultOf2(rslt.Error[string](err), rslt.Value("input"))

		assert.True(t, result.IsErr())
		assert.ErrorIs(t, result.Error(), err)
	})

	t.Run("InputError", func(t *testing.T) {
		err := errors.New("input err")
		result := ResultOf2(rslt.Value("cmd"), rslt.Error[string](err))

		assert.True(t, result.IsErr())
		assert.ErrorIs(t, result.Error(), err)
	})
}
