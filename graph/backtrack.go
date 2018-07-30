package graph

import ()

func (g *Graph) Findk(node *Node, k int) bool {
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	return g.findkUtil(node, visited, 0, k)
}

func (g *Graph) findkUtil(node *Node, visited map[*Node]bool, index int, k int) bool {
	if index >= k {
		return true
	}
	visited[node] = true
	flag := false
	for _, e := range node.edges {
		if !visited[e.dest] {
			flag = flag || g.findkUtil(e.dest, visited, index+1, k)
		}
	}
	visited[node] = false
	return flag
}
