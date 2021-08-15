package datastructures

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestHeap(t *testing.T) {
	g := Goblin(t)

	g.Describe("Heap", func() {
		g.It("Creating a heap", func() {
			h := Heap{}
			h.Init([]int{2, 4, 8, 9, 7, 10, 9, 15, 20, 13})
			// 	       2
			// 	   /      \
			//        4         8
			//     /    \      / \
			//    9      7   10   9
			//   / \    /
			//  15 20  13
			g.Assert(h.hasParent(0)).Equal(false)
			g.Assert(h.parent(1)).Equal(2)
			g.Assert(h.parent(2)).Equal(2)
			g.Assert(h.parent(3)).Equal(4)
			g.Assert(h.parent(4)).Equal(4)
			g.Assert(h.parent(5)).Equal(8)
			g.Assert(h.parent(6)).Equal(8)
			g.Assert(h.parent(7)).Equal(9)
			g.Assert(h.parent(8)).Equal(9)
			g.Assert(h.parent(9)).Equal(7)
			g.Assert(h.left(3)).Equal(15)
			g.Assert(h.right(3)).Equal(20)
			g.Assert(h.hasLeft(4)).Equal(true)
			g.Assert(h.hasRight(4)).Equal(false)
			g.Assert(h.hasLeft(7)).Equal(false)
			g.Assert(h.hasRight(7)).Equal(false)

			h.add(3) // 3 will be added as right child of 7, and it should bubble up to the right position below 2.
			// 	       2
			// 	   /      \
			//        3         8
			//     /    \      / \
			//    9      4   10   9
			//   / \    / \
			//  15 20  13  7
			g.Assert(h.left(0)).Equal(3)
			g.Assert(h.left(1)).Equal(9)
			g.Assert(h.right(1)).Equal(4)
			g.Assert(h.left(4)).Equal(13)
			g.Assert(h.right(4)).Equal(7)

			h.Poll() // the root element (2) will be deleted and the last element (7) will be added on its place and should be bubbling down to the right position.
			// 	       3
			// 	   /      \
			//        4         8
			//     /    \      / \
			//    9      7   10   9
			//   / \    /
			//  15 20  13
			g.Assert(h.parent(1)).Equal(3)
			g.Assert(h.parent(2)).Equal(3)
			g.Assert(h.parent(3)).Equal(4)
			g.Assert(h.parent(4)).Equal(4)
			g.Assert(h.parent(5)).Equal(8)
			g.Assert(h.parent(6)).Equal(8)
			g.Assert(h.parent(7)).Equal(9)
			g.Assert(h.parent(8)).Equal(9)
			g.Assert(h.parent(9)).Equal(7)

			h.Poll() // the root element (3) will be deleted and the last element (13) will be added on its place and should be bubbling down to the right position.
			// 	       4
			// 	   /      \
			//        7         8
			//     /    \      / \
			//    9      13   10   9
			//   / \
			//  15 20
			g.Assert(h.parent(1)).Equal(4)
			g.Assert(h.parent(2)).Equal(4)
			g.Assert(h.parent(3)).Equal(7)
			g.Assert(h.parent(4)).Equal(7)
			g.Assert(h.parent(5)).Equal(8)
			g.Assert(h.parent(6)).Equal(8)
			g.Assert(h.parent(7)).Equal(9)
			g.Assert(h.parent(8)).Equal(9)
			g.Assert(h.hasLeft(9)).Equal(false)
			g.Assert(h.hasRight(9)).Equal(false)
		})
	})

}
