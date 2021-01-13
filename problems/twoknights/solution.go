package twoknights

import (
	"fmt"
)

func TwoKnights(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println((i * i * (i*i - 1) / 2) - (8 * (i - 2) * (i - 1) / 2))
	}
}
