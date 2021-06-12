package repetitions_test

import (
	"testing"

	. "github.com/jn-lp/go-algorithms/problems/repetitions"
	"github.com/stretchr/testify/assert"
)

func TestRepetitions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str         string
		repetitions int
	}{{
		str:         "ATTCGGGA",
		repetitions: 3,
	}}

	for _, test := range tests {
		assert.Equal(
			t,
			test.repetitions,
			Repetitions(test.str),
		)
	}
}
