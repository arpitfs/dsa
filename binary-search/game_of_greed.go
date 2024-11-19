package binarysearch

import "fmt"

func greedyCoins() {
	coins := []int{1, 2, 3, 4}
	partitions := 3
	start := 0
	end := 0
	ans := 0
	for _, val := range coins {
		end += val
	}

	for start <= end {
		mid := (start + end) / 2
		isPossible := divideIntoPartitions(coins, len(coins), partitions, mid)

		if isPossible {
			start = mid + 1
			ans = mid
		} else {
			end = mid - 1
		}
	}

	fmt.Println(ans)
}

func divideIntoPartitions(coins []int, n, p, min_coins int) bool {
	partitions := 0
	current_fried := 0
	for i := 0; i < n; i++ {
		if current_fried+coins[i] >= min_coins {
			partitions += 1
			current_fried = 0
		} else {
			current_fried += coins[i]
		}
	}

	return partitions >= p
}
