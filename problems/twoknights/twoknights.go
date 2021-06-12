package twoknights

func TwoKnights(n int) (ways []int) {
	for i := 1; i <= n; i++ {
		ways = append(ways, (i*i*(i*i-1)/2)-(8*(i-2)*(i-1)/2))
	}

	return ways
}
