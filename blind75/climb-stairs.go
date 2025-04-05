package blind75

import "fmt"

func climb() {
	input := 5
	fmt.Println(climbStairs(input))
}

func climbStairs(n int) int {
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
