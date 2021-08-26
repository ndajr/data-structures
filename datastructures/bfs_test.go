package datastructures

import (
	"math"
	"testing"

	. "github.com/franela/goblin"
)

func TestBreadthFirstSearch(t *testing.T) {
	g := Goblin(t)

	g.Describe("BreadthFirstSearch", func() {
		g.It("Should solve the Mango seller problem BFS algorithm", func() {
			// Problem from Grokking algorithms book
			// Suppose you’re the proud owner of a mango farm. You’re looking for a mango seller who can sell your mangoes.
			// Find the shortest path to the mango seller, first checking first-degree connections, then second-degree connections and so on. Check wether the person's name ends with the letter m, if it does they are a mango seller.
			findMangoSeller := func(source string, connections map[string][]string) string {
				vis := map[string]bool{}
				queue := Queue{}
				queue.Enqueue(Node{data: source})
				for !queue.isEmpty() {
					curr := queue.Dequeue().data.(string)
					if vis[curr] {
						continue
					}
					vis[curr] = true

					for _, neighbor := range connections[curr] {
						isMangoSeller := neighbor[len(neighbor)-1:] == "m"
						if isMangoSeller {
							return neighbor
						} else {
							queue.Enqueue(Node{data: neighbor})
						}
					}
				}
				return ""
			}

			g.Assert(findMangoSeller("you", map[string][]string{
				"you":    {"alice", "bob", "claire"},
				"bob":    {"anuj", "peggy"},
				"alice":  {"peggy"},
				"claire": {"thom", "jonny"},
				"anuj":   {},
				"peggy":  {},
				"thom":   {},
				"jonny":  {},
			})).Equal("thom")

		})

		g.It("Should solve the Wizzard meetup problem using BFS algorithm", func() {
			// Problem - Wizard Meet-up
			// There are 10 wizards at a meet-up. Each wizard has level 0-9(the index of the input) and knows a few other wizard here.
			// Find the cheapest way for wizard level X to meet wizard level Y.
			// Introductions have a cost that equal the square of the difference in levels.
			getMinCostPath := func(source int, target int, wizzards [][]int) ([]int, float64) {
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
