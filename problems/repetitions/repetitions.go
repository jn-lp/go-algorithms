package repetitions

import (
	"strings"
)

func Repetitions(str string) (highest int) {
	s := strings.Split(str, "")
	current := 1

	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if current > highest {
				highest = current
			}

			current = 0
		}
		current++
	}

	return highest
}
