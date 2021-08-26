package datastructures

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestLinkedList(t *testing.T) {
	g := Goblin(t)

	g.Describe("Node", func() {
		g.It("Should create linked Nodes with string", func() {
			n2 := Node{next: nil, data: "node2"}
			n1 := Node{next: nil, data: "node1"}
			g.Assert(n1.data).Equal("node1")
			g.Assert(n1.next == nil).Equal(true)
			n1.next = &n2
			g.Assert(n1.next.data).Equal("node2")
		})

		g.It("Should create linked Nodes with int", func() {
			n2 := Node{next: nil, data: 2}
			n1 := Node{next: nil, data: 1}
			g.Assert(n1.data).Equal(1)
			g.Assert(n1.next == nil).Equal(true)
			n1.next = &n2
			g.Assert(n1.next.data).Equal(2)
		})
	})

	g.Describe("LinkedList", func() {
		g.It("Should create a LinkedList", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			l := LinkedList{}
			l.Append(n1)
			l.Append(n2)
			g.Assert(l.Len()).Equal(2)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node1 node2]")

			n0 := Node{next: nil, data: "node0"}
			l.Prepend(n0)
			g.Assert(l.Len()).Equal(3)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node0 node1 node2]")
		})

		g.It("Should delete a node from LinkedList", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			n3 := Node{next: nil, data: "node3"}
			l := LinkedList{}
			l.Append(n1)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node1]")
			l.Delete(n1)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[]")

			l.Append(n1)
			l.Append(n2)
			l.Append(n3)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node1 node2 node3]")
			l.Delete(n2)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node1 node3]")
		})

		g.It("Should find a node from LinkedList", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			n3 := Node{next: nil, data: "node3"}
			unknown := Node{next: nil, data: "unknown node"}
			l := LinkedList{}
			l.Append(n1)
			l.Append(n2)
			l.Append(n3)
			g.Assert(l.Find(n3)).Equal(2)
			g.Assert(l.Find(unknown)).Equal(-1)
		})

		g.It("Should reverse the order of a LinkedList", func() {
			n1 := Node{next: nil, data: "node1"}
			n2 := Node{next: nil, data: "node2"}
			n3 := Node{next: nil, data: "node3"}
			l := LinkedList{}
			l.Append(n1)
			l.Append(n2)
			l.Append(n3)
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node1 node2 node3]")
			l.Reverse()
			g.Assert(fmt.Sprintf("%s", l)).Equal("[node3 node2 node1]")
		})
	})

}
