package increasingarray

import (
	"fmt"
)

func IncreasingArray(n int, integers []int) {
	maxdif := 0
	for i := 1; i < n; i++ {
		if prev := integers[i-1] + 1; integers[i] < prev {
			maxdif, integers[i] = prev-integers[i], prev
		}
	}

	fmt.Println(maxdif)
}
