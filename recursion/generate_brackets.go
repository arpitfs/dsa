package recursion

import "fmt"

func GenerateBrackets() {
	generate("", 3, 0, 0, 0)
}

func generate(output string, n, open, close, i int) {
	if i == 2*n {
		fmt.Println(output)
		return
	}

	if open < n {
		generate(output+"(", n, open+1, close, i+1)
	}

	if close < open {
		generate(output+")", n, open, close+1, i+1)
	}
}
