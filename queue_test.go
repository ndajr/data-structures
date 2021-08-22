package datastructures

import (
	"fmt"
	"math"
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

func getMinCostPath(source int, target int, wizzards [][]int) ([]int, float64) {
	n := len(wizzards)
	prev := make([]int, n)
	dist := make([]float64, n)
	vis := map[int]bool{}
	for i := 0; i < n; i++ {
		// Initialize prev as -1 since the default value (0) is the the first node which would be wrong in this case.
		prev[i] = -1
		// Initialize distance slice with Infinity for all elements. If the target remains Inifinity at the end, it means the target is unreachable and it is not possible to find a path on the graph.
		dist[i] = math.Inf(1)
	}
	dist[source] = 0
	queue := Queue{}
	queue.Enqueue(Node{data: source})

	for !queue.isEmpty() {
		curr := queue.Dequeue().data.(int)
		neighbors := wizzards[curr]
		if vis[curr] {
			continue
		}
		vis[curr] = true
		for _, neighbor := range neighbors {
			weight := math.Pow(float64(neighbor-curr), 2)
			if dist[curr]+weight < dist[neighbor] {
				prev[neighbor] = curr
				dist[neighbor] = dist[curr] + weight
			}
			queue.Enqueue(Node{data: neighbor})
		}
	}

	var path []int
	var cost float64
	if dist[target] == math.Inf(1) {
		return path, cost
	}
	for t := target; t != -1; t = prev[t] {
		cost += dist[t]
		path = append(path, t)
	}
	path = reverse(path)
	return path, cost
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

	g.Describe("Breadth-first search", func() {
		g.It("Should traverse a graph using BFS algorithm and find the shortest path using a Queue", func() {
			// Problem - Wizard Meet-up
			// There are 10 wizards at a meet-up. Each wizard has level 0-9(the index of the input) and knows a few other wizard here.
			// Find the cheapest way for wizard level X to meet wizard level Y.
			// Introductions have a cost that equal the square of the difference in levels.
			path, cost := getMinCostPath(0, 9, [][]int{
				{1, 5, 9}, {2, 3, 9}, {4}, {}, {}, {9}, {}, {}, {}, {},
			})

			g.Assert(path).Equal([]int{0, 5, 9})
			g.Assert(int(cost)).Equal(66)

			path, cost = getMinCostPath(1, 9, [][]int{
				{1, 5, 2}, {8, 6, 4}, {7, 9, 3}, {8, 1}, {9}, {8, 7}, {4, 6}, {9, 4}, {1}, {1, 4},
			})
			g.Assert(path).Equal([]int{1, 4, 9})
			g.Assert(int(cost)).Equal(43)
		})
	})

}
