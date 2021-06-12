package weirdalgorithm_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/weirdalgorithm"
	"github.com/stretchr/testify/assert"
)

func TestWeirdAlgorithm(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n      int
		result []int
	}{{
		n:      3,
		result: []int{3, 10, 5, 16, 8, 4, 2, 1},
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.result,
			WeirdAlgorithm(test.n),
		)
	}
}
