package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_Contains(t *testing.T) {

	set := make(Set[int])

	assert.False(t, set.Contains(3))

	set.Add(3)

	assert.True(t, set.Contains(3))
}
