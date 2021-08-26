package datastructures

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func reverse(s []int) []int {
	a := make([]int, len(s))
	copy(a, s)
	// From https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}

func TestQueue(t *testing.T) {
	g := Goblin(t)

	g.Describe("Queue", func() {
		g.It("Should enqueue and dequeue elements from the queue", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			s := Queue{}
			s.Enqueue(n1)
			s.Enqueue(n2)
			g.Assert(fmt.Sprintf("%s", s)).Equal("[node1 node2]")
			s.Dequeue()
			g.Assert(fmt.Sprintf("%s", s)).Equal("[node2]")
			s.Dequeue()
			g.Assert(fmt.Sprintf("%s", s)).Equal("[]")
		})
	})

}
