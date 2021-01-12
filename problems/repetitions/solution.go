package repetitions

import (
	"fmt"
	"strings"
)

func Repetitions(str string) {
	s := strings.Split(str, "")
	highest, current := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if current > highest {
				highest = current
			}
			current = 0
		}
		current++
	}

	fmt.Println(highest)
}
