package weirdalgorithm

func WeirdAlgorithm(n int, results ...int) []int {
	results = append(results, n)

	if n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		return WeirdAlgorithm(n, results...)
	}

	return results
}
