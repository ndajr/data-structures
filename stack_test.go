package datastructures

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestStack(t *testing.T) {
	g := Goblin(t)

	g.Describe("Stack", func() {
		g.It("Should push and pop elements from the stack", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			s := Stack{}
			s.Push(n1)
			s.Push(n2)
			g.Assert(fmt.Sprintf("%s", s)).Equal("[node1 node2]")
			s.Pop()
			g.Assert(fmt.Sprintf("%s", s)).Equal("[node1]")
			s.Pop()
			g.Assert(fmt.Sprintf("%s", s)).Equal("[]")
		})
	})

}
