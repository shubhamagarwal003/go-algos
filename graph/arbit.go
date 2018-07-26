package graph

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
