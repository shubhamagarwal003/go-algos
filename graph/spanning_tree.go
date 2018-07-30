package graph

import (
	"fmt"
	"math"
	"sort"
)

type Edges []*Edge

func (e Edges) Len() int           { return len(e) }
func (e Edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Edges) Less(i, j int) bool { return e[i].weight < e[j].weight }

func (g *Graph) Prims() {
	edges := g.primsUtil(g.nodes[0])
	for _, edge := range edges {
		fmt.Println(edge.src.value, edge.dest.value, edge.weight)
	}
}

func (g *Graph) primsUtil(node *Node) []*Edge {
	nodes := []*Node{node}
	included := make(map[*Node]bool)
	for _, n := range g.nodes {
		included[n] = false
	}
	included[node] = true
	edges := []*Edge{}
	for len(nodes) < len(g.nodes) {
		minValue := math.MaxInt32
		minEdge := &Edge{}
		for _, n := range nodes {
			for _, e := range n.edges {
				if e.weight <= minValue && !included[e.dest] {
					minEdge = e
					minValue = e.weight
				}
			}
		}
		included[minEdge.dest] = true
		nodes = append(nodes, minEdge.dest)
		edges = append(edges, minEdge)
	}
	return edges
}

func (g *Graph) Kruskal() {
	edges := g.kruskalUtil()
	for _, edge := range edges {
		fmt.Println(edge.src.value, edge.dest.value, edge.weight)
	}
}

// This is a modified version in which I am keeping track of included nodes.
// That can be avoided by checking whether included node doesn't form a cycle.
func (g *Graph) kruskalUtil() []*Edge {
	edges := []*Edge{}
	includedEdges := []*Edge{}
	for _, n := range g.nodes {
		edges = append(edges, n.edges...)
	}
	sort.Sort(Edges(edges))
	for len(includedEdges) < len(g.nodes)-1 {
		// fmt.Println(len(includedEdges), len(edges))
		for _, e := range edges {
			if !g.cycle(includedEdges, e) {
				includedEdges = append(includedEdges, e)
			}
		}
	}
	return includedEdges
}

func (g *Graph) cycle(edges []*Edge, edge *Edge) bool {
	node1 := edge.src
	node2 := edge.dest
	visited := make(map[*Node]bool)
	for _, n := range g.nodes {
		visited[n] = false
	}
	return g.cycleUtil(edges, node1, node2, visited)
}

func (g *Graph) cycleUtil(edges []*Edge, node *Node, dest *Node, visited map[*Node]bool) bool {
	if node == dest {
		return true
	}
	flag := false
	for _, e := range edges {
		if e.dest == node && !visited[e.src] {
			visited[e.src] = true
			flag = flag || g.cycleUtil(edges, e.src, dest, visited)
		}
		if e.src == node && !visited[e.dest] {
			visited[e.dest] = true
			flag = flag || g.cycleUtil(edges, e.dest, dest, visited)
		}
	}
	return flag
}
