package missingnumber

import "fmt"

func MissingNumber(n int, numbers []int) {
	n = n * (n + 1) / 2
	for _, number := range numbers {
		n -= number
	}

	fmt.Println(n)
}
