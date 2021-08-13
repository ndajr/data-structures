package datastructures

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestSelectionSort(t *testing.T) {
	g := Goblin(t)

	g.Describe("SelectionSort", func() {
		g.It("Should return ordered the given array", func() {
			g.Assert(SelectionSort([]int{5, 3, 6, 2, 10})).Equal([]int{2, 3, 5, 6, 10})
		})
	})

	g.Describe("BubbleSort", func() {
		g.It("Should return ordered the given array", func() {
			g.Assert(BubbleSort([]int{5, 3, 6, 2, 10})).Equal([]int{2, 3, 5, 6, 10})
		})
	})

}
