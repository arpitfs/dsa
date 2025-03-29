package practice

import "fmt"

type Stack struct {
	data []string
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (bool, string) {
	if len(s.data) == 0 {
		return false, ""
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return true, item
}
func CreateStack() *Stack {
	return &Stack{}
}

func checkParanthesis() {
	fmt.Println(validParanthesis())
}

func validParanthesis() bool {
	input := "[{]"
	set := make(map[string]string)

	set["]"] = "["
	set["}"] = "{"
	stack := CreateStack()
	fmt.Println(stack)

	for _, value := range input {
		if value == '[' || value == '{' {
			stack.Push(string(value))
		} else {
			if len(stack.data) == 0 {
				return false
			}
			_, pair := stack.Pop()
			if pair != set[string(value)] {
				return false
			}
		}
	}

	return len(stack.data) == 0
}
