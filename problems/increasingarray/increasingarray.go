package increasingarray

func IncreasingArray(n int, integers []int) (maxDiff int) {
	for i := 1; i < n; i++ {
		if prev := integers[i-1] + 1; integers[i] < prev {
			maxDiff, integers[i] = prev-integers[i], prev
		}
	}

	return maxDiff
}
