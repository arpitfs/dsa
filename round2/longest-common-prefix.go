package round2

import (
	"fmt"
	"strings"
)

func longestCommonPrefix() {
	input := []string{"flower", "flow", "flight"}
	prefix := input[0]

	for i := 1; i < len(input); i++ {
		for !strings.HasPrefix(input[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				fmt.Println("")
				break
			}
		}
	}
	fmt.Println(prefix)
}
