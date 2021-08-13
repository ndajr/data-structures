package datastructures

import "fmt"

type Stack struct {
	list LinkedList
}

func (s *Stack) Push(n Node) {
	s.list.Append(n)
}

func (s Stack) isEmpty() bool {
	return s.list.Head == nil
}

func (s *Stack) Pop() bool {
	if s.isEmpty() {
		return false
	}
	s.list.Delete(*s.list.Tail)
	return true
}

func (s Stack) String() string {
	return fmt.Sprint(s.list)
}
