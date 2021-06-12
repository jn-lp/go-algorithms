package twosets_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/twosets"
	"github.com/stretchr/testify/assert"
)

func TestRepetitions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		integer  int
		sets     [2][]int
		possible bool
	}{{
		integer:  7,
		sets:     [2][]int{{7, 6, 1}, {5, 4, 3, 2}},
		possible: true,
	}, {
		integer:  6,
		possible: false,
	}}

	for _, test := range tests {
		sets, possible := TwoSets(test.integer)

		assert.Equal(
			t,
			test.possible,
			possible,
		)

		assert.Equal(
			t,
			test.sets,
			sets,
		)
	}
}
