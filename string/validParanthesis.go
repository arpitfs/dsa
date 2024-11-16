package string

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func paranthesis() {
	input := "{)[]"
	fmt.Println(checkValidParanthesis(input))

}

func checkValidParanthesis(input string) bool {
	tracker := make(map[rune]rune)
	tracker[')'] = '('
	tracker[']'] = '['
	tracker['}'] = '{'

	tray := stack.New()
	for i := 0; i < len(input); i++ {
		if _, e := tracker[rune(input[i])]; !e {
			tray.Push(rune(input[i]))
		} else {
			if tray.Len() > 0 {
				// input[i] = )
				valueFromTray := tray.Pop().(rune) // (
				if valueFromTray != tracker[rune(input[i])] {
					return false
				}
			} else {
				return false
			}
		}
	}

	if tray.Len() > 0 {
		return false
	}

	return true
}
