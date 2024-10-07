// package main

// import "fmt"

// func main() {
// 	input := []int{2, 1, 3, 5, 8}
// 	target := 9

// 	set := make(map[int]int)

// 	for i := 0; i <= len(input)-1; i++ {

// 		data := target - input[i]

// 		if _, exists := set[data]; exists {
// 			fmt.Println(i, set[data])
// 		} else {
// 			set[i] = input[i]
// 		}

// 	}
// }
