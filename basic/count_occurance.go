package basic

import "fmt"

func countOccurance() {
	input := "programming"
	letter := 'g'
	count := 0
	for _, val := range input {
		if val == letter {
			count++
		}
	}
	fmt.Println(count)
}
