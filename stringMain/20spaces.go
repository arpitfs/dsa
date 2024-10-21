package stringMain

import (
	"fmt"
	"strings"
)

func spaces() {
	input := "hello world, how are you?"
	fmt.Println(len(input))
	appender := "%20"
	var builder strings.Builder

	for _, char := range input {
		if char == ' ' {
			builder.WriteString(appender)
		} else {
			builder.WriteString(string(char))
		}
	}

	fmt.Println(builder.String())
}
