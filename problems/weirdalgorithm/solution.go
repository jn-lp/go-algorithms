package weirdalgorithm

import "fmt"

func WeirdAlgorithm(n int) {
	fmt.Printf("%d ", n)
	if n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		WeirdAlgorithm(n)
	}
}
