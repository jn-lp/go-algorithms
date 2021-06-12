package numberspiral_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/numberspiral"
	"github.com/stretchr/testify/assert"
)

func TestNumberSpiral(t *testing.T) {
	t.Parallel()

	tests := []struct {
		n          int
		coordsSets [][2]float64
		spiral     []float64
	}{{
		n:          5,
		coordsSets: [][2]float64{{2, 3}, {1, 1}, {4, 2}},
		spiral:     []float64{8, 1, 15},
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.spiral,
			NumberSpiral(test.n, test.coordsSets...),
		)
	}
}
