package basic

import "fmt"

func removeOccurance() {
	input := "hello world"
	letter := 'l'
	result := ""
	for _, val := range input {
		if val != letter {
			result += string(val)
		}
	}
	fmt.Println(result)

}
