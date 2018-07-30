package graph

import (
	"fmt"
)

func (g *Graph) CountAllPaths(node1 *Node, node2 *Node) int {
	return g.countAllPathsUtil(node1, node2)
}

func (g *Graph) countAllPathsUtil(node1 *Node, node2 *Node) int {
	sum := 0
	for _, e := range node1.edges {
		if e.dest == node2 {
			sum++
		} else {
			sum += g.countAllPathsUtil(e.dest, node2)
		}
	}
	return sum
}

func (g *Graph) CycleLength(n int) {
	for _, node := range g.nodes {
		g.cycleLengthUtils(node, node, nil, n)
	}
}

func (g *Graph) cycleLengthUtils(source *Node, node *Node, parent *Node, n int) {
	for i := 1; i < n; i++ {
		for _, e := range node.edges {
			if e.dest != parent {
				g.cycleLengthUtils(source, e.dest, node, n-i)
			}
		}
	}
	for _, e := range node.edges {
		if e.dest == source && e.dest != parent {
			fmt.Println("N node Cycle Found", node.value, e.dest.value)
		}
	}
}
