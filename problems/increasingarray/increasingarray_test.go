package increasingarray_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/increasingarray"
	"github.com/stretchr/testify/assert"
)

func TestIncreasingArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n        int
		integers []int
		maxDiff  int
	}{{
		n:        5,
		integers: []int{3, 2, 5, 1, 7},
		maxDiff:  5,
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.maxDiff,
			IncreasingArray(test.n, test.integers),
		)
	}
}
