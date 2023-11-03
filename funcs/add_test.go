package funcs_test

import (
	"testing"

	"github.com/izaakdale/ci-test/funcs"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	for i := -10; i <= 10; i++ {
		assert.Equal(t, i+i, funcs.Add(i, i))
	}
}
