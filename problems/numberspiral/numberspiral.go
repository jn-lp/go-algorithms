package numberspiral

import (
	"math"
)

func NumberSpiral(n int, coordsSet ...[2]float64) (spiral []float64) {
	for _, coords := range coordsSet {
		var m float64
		if coords[0] > coords[1] {
			m = coords[0]
		} else {
			m = coords[1]
		}

		spiral = append(spiral, (coords[0]-coords[1])*math.Pow(-1, m)+m*m-m+1)
	}

	return spiral
}
