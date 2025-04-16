package round2

// import (
// 	"fmt"
// 	"sort"
// )

// type Pair struct {
// 	value int
// 	index int
// }

// func minSwaps() {
// 	arr := []int{4, 3, 2, 1}
// 	fmt.Println("Minimum swaps required:", minimumSwaps(arr))
// }

// func minimumSwaps(arr []int) int {
// 	n := len(arr)
// 	pairs := make([]Pair, n)

// 	for i := 0; i < n; i++ {
// 		pairs[i] = Pair{arr[i], i}
// 	}

// 	sort.Slice(pairs, func(i, j int) bool {
// 		return pairs[i].value < pairs[j].value
// 	})

// 	visited := make([]bool, n)
// 	swaps := 0

// 	for i := 0; i < n; i++ {
// 		if visited[i] || pairs[i].index == i {
// 			continue
// 		}
// 	}
// }
