package datastructures

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestBinarySearch(t *testing.T) {
	g := Goblin(t)

	g.Describe("BinarySearch", func() {
		g.It("Should return -1 if not found", func() {
			g.Assert(BinarySearch([]int{1, 3, 5, 7, 9}, -1)).Equal(-1)
		})
		g.It("Should return the index if found", func() {
			g.Assert(BinarySearch([]int{100, 200, 300, 500, 900}, 900)).Equal(4)
			g.Assert(BinarySearch([]int{1, 3, 5, 7, 9}, 3)).Equal(1)
		})
	})

}
