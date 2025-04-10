package round2

import "fmt"

func minRotatedArray() {
	input := []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	ans := findMinInRotatedArray(input)
	fmt.Println(ans)
}

func findMinInRotatedArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]

}
