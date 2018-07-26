package graph

import (
// "fmt"
)

func (g *Graph) DetectCycleDirected() bool {
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	cycle := false
	currentNodes := []*Node{}
	for _, node := range g.nodes {
		if !visited[node] {
			cycle = cycle || g.detectCycleDirectedUtil(node, visited, currentNodes)
		}
	}
	return cycle
}

func (g *Graph) detectCycleDirectedUtil(node *Node, visited map[*Node]bool, currentNodes []*Node) bool {
	visited[node] = true
	currentNodes = append(currentNodes, node)
	flag := false
	for _, e := range node.edges {
		if !visited[e.dest] {
			visited[e.dest] = true
			flag = flag || g.detectCycleDirectedUtil(e.dest, visited, currentNodes)
		} else {
			for _, n := range currentNodes {
				// fmt.Println(n.value, node.value, n, node)
				if n == e.dest {
					return true
				}
			}
		}
	}
	i := 0
	for i < len(currentNodes) {
		if currentNodes[i] == node {
			break
		}
		i++
	}
	currentNodes = append(currentNodes[:i], currentNodes[i+1:]...)
	return flag
}

func (g *Graph) DetectCycleUndirected() bool {
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	cycle := false
	currentNodes := []*Node{}
	for _, node := range g.nodes {
		if !visited[node] {
			cycle = cycle || g.detectCycleUndirectedUtil(node, visited, currentNodes, nil)
		}
	}
	return cycle
}

func (g *Graph) detectCycleUndirectedUtil(node *Node, visited map[*Node]bool, currentNodes []*Node, parent *Node) bool {
	visited[node] = true
	currentNodes = append(currentNodes, node)
	flag := false
	for _, e := range node.edges {
		if !visited[e.dest] {
			visited[e.dest] = true
			flag = flag || g.detectCycleUndirectedUtil(e.dest, visited, currentNodes, node)
		} else {
			for _, n := range currentNodes {
				// fmt.Println(n.value, node.value, n, node)
				if n == e.dest && parent != e.dest {
					return true
				}
			}
		}
	}
	i := 0
	for i < len(currentNodes) {
		if currentNodes[i] == node {
			break
		}
		i++
	}
	currentNodes = append(currentNodes[:i], currentNodes[i+1:]...)
	return flag
}
