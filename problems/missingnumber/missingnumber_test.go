package missingnumber_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/missingnumber"
	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n       int
		numbers []int
		missing int
	}{{
		n:       5,
		numbers: []int{2, 3, 1, 5},
		missing: 4,
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.missing,
			MissingNumber(test.n, test.numbers),
		)
	}
}
