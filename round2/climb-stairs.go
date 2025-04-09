package round2

import "fmt"

func climbStairs() {
	n := 5
	fmt.Println(climbPossible(n))
}

func climbPossible(n int) int {
	if n <= 2 {
		return n
	}

	f := 1
	s := 2
	for i := 3; i <= n; i++ {
		t := f + s
		f = s
		s = t
	}

	return s
}
