package round2

import "fmt"

func isValidParanthesis() {
	input := "({)"
	fmt.Println(isValid(input))
}

type Stack struct {
	data []string
}

func CreateStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (bool, string) {
	if len(s.data) <= 0 {
		return false, ""
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return true, item
}

func isValid(s string) bool {
	checker := make(map[rune]rune)
	checker[']'] = '['
	checker['}'] = '{'
	checker[')'] = '('

	stack := CreateStack()
	for _, v := range s {
		if v == '[' || v == '{' || v == '(' {
			stack.Push(string(v))
		} else {
			ok, item := stack.Pop()
			if !ok {
				return false
			} else {
				if item != string(checker[v]) {
					return false
				}
			}
		}
	}

	return len(stack.data) == 0
}
