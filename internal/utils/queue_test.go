package utils_test

import (
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("Length", func(t *testing.T) {

		q := utils.Queue[string]{}

		assert.Equal(t, 0, q.Len())

		q.Add("foo")

		assert.Equal(t, 1, q.Len())

		q.Add("foo")
		q.Add("bar")

		assert.Equal(t, 3, q.Len())

		q.RemoveFirst()

		assert.Equal(t, 2, q.Len())
	})

	t.Run("Remove", func(t *testing.T) {
		q := utils.Queue[string]{}

		q.Add("foo")
		q.Add("bar")
		q.Add("bar")
		q.Add("baz")

		assert.Equal(t, "foo", q.RemoveFirst())
		assert.Equal(t, "bar", q.RemoveFirst())
		assert.Equal(t, "bar", q.RemoveFirst())
		assert.Equal(t, "baz", q.RemoveFirst())
	})

}
