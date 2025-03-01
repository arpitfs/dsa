package array

import "fmt"

// Check if a Pair Exists with Given Sum in a Sorted Array

func checkPairExists() {
	input := []int{1, 2, 3, 4, 6}
	target := 9
	fmt.Println(isExists(input, target))
	fmt.Println(isExistsUsingTwoPointer(input, target))

}

func isExists(input []int, target int) bool {
	set := make(map[int]bool)

	for _, val := range input {
		set[val] = true
	}

	for i := 0; i < len(input); i++ {
		getValue := target - input[i]
		if set[getValue] {
			return true
		}
	}

	return false
}

// Using two pointer

func isExistsUsingTwoPointer(input []int, target int) bool {
	left := 0
	right := len(input) - 1
	for left < right {
		if input[left]+input[right] == target {
			return true
		} else if input[left]+input[right] > target {
			right--
		} else {
			left++
		}
	}
	return false
}
