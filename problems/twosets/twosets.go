package twosets

func TwoSets(n int) (sets [2][]int, possible bool) {
	total := n * (n + 1) / 2
	if total&1 == 1 {
		return sets, false
	}

	total /= 2

	for i := n; i > 0; i-- {
		if i <= total {
			total -= i
			sets[0] = append(sets[0], i)
		} else {
			sets[1] = append(sets[1], i)
		}
	}

	if total != 0 {
		sets[0] = append(sets[0], total)
	}

	return sets, true
}
