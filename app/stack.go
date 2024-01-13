package app

import "fmt"

func StackMain() {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Pop()) // 1

	fmt.Println(stack.IsEmpty()) // true
	fmt.Println(stack.Size())    // 0
}

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0),
	}
}

func (s *Stack) Push(data interface{}) {
	s.data = append(s.data, data)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return data
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}
