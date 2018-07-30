package graph

import (
	"fmt"
	"math"
)

func (g *Graph) BellmanFord(source *Node) map[*Node]int {
	cost := make(map[*Node]int)
	for _, node := range g.nodes {
		cost[node] = math.MaxInt32 //Very
	}
	cost[source] = 0
	for i := 0; i < len(g.nodes); i++ {
		for _, node := range g.nodes {
			for _, e := range node.edges {
				if cost[e.dest] > cost[node]+e.weight {
					cost[e.dest] = cost[node] + e.weight
				}
			}
		}
	}
	for k, v := range cost {
		fmt.Println(k.value, v)
	}
	return cost
}
