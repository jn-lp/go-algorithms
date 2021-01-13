package twosets

import "fmt"

func TwoSets(n int) {
	total := n * (n + 1) / 2
	if total&1 == 1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
		//TODO
	}
}
