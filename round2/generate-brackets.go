package round2

import "fmt"

func generateBrackets() {
	generate("", 3, 0, 0, 0)
}

func generate(output string, input, open, close, index int) {
	if index == 2*input {
		fmt.Println(output)
		return
	}

	if open < input {
		generate(output+"(", input, open+1, close, index+1)
	}

	if close < open {
		generate(output+")", input, open, close+1, index+1)
	}
}
