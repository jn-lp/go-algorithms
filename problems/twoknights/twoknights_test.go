package twoknights_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/twoknights"
	"github.com/stretchr/testify/assert"
)

func TestWeirdAlgorithm(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n    int
		ways []int
	}{{
		n:    8,
		ways: []int{0, 6, 28, 96, 252, 550, 1056, 1848},
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.ways,
			TwoKnights(test.n),
		)
	}
}
