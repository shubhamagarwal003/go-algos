package graph

import (
	"fmt"
	"math"
)

func (g *Graph) Dijkstra(src *Node) {
	weight := make(map[*Node]int)
	weightEdge := make(map[*Node]*Edge)
	for _, n := range g.nodes {
		weight[n] = math.MaxInt32
		weightEdge[n] = &Edge{}
	}
	weight[src] = 0
	edges := g.dijkstraUtil([]*Node{}, []*Edge{}, weight, weightEdge)
	for _, e := range edges {
		if e != nil && e.src != nil && e.dest != nil {
			fmt.Println(e.src.value, e.dest.value, e.weight)
		}
	}
}

func (g *Graph) dijkstraUtil(nodes []*Node, edges []*Edge, weight map[*Node]int, weightEdge map[*Node]*Edge) []*Edge {
	if len(nodes) == len(g.nodes) {
		return edges
	}
	minValue := math.MaxInt32
	minNode := &Node{}
	for k, v := range weight {
		flag := false
		for _, n := range nodes {
			if k == n {
				flag = true
			}
		}
		if !flag {
			if v < minValue {
				minValue = v
				minNode = k
			}
		}
	}
	nodes = append(nodes, minNode)
	edges = append(edges, weightEdge[minNode])
	for _, e := range minNode.edges {
		if weight[e.dest] > minValue+e.weight {
			weight[e.dest] = minValue + e.weight
			weightEdge[e.dest] = e
		}
	}
	return g.dijkstraUtil(nodes, edges, weight, weightEdge)
}
