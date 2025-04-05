package blind75

import "fmt"

func jump() {
	nums := []int{2, 3, c, 1, 4}
	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	maxJumps := 0
	for i, v := range nums {
		if i > maxJumps {
			return false
		}

		if i+v > maxJumps {
			maxJumps = i + v
		}
	}
	return true
}
