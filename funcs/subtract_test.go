package funcs_test

import (
	"testing"

	"github.com/izaakdale/ci-test/funcs"
	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	for i := -10; i <= 10; i++ {
		j := 7
		assert.Equal(t, j-i, funcs.Subtract(j, i))
	}
}
