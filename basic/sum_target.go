package basic

import (
	"fmt"
)

func sumTarget() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(getSumIndex(nums, target))
}

func getSumIndex(nums []int, target int) (int, int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		currentSum := nums[left] + nums[right]
		if currentSum == target {
			return left, right
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return -1, -1
}
