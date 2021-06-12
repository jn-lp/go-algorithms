package missingnumber

func MissingNumber(n int, numbers []int) (missing int) {
	missing = n * (n + 1) / 2
	for _, number := range numbers {
		missing -= number
	}

	return missing
}
