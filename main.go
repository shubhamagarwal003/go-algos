package main

import (
	"fmt"
	"github.com/shubhamagarwal003/go-algos/graph"
	// "reflect"
)

func main() {
	g := graph.NewGraph()
	n1 := graph.NewNode(1)
	n2 := graph.NewNode(2)
	n3 := graph.NewNode(3)
	n4 := graph.NewNode(4)
	n5 := graph.NewNode(5)
	n6 := graph.NewNode(6)
	// n7 := graph.NewNode(7)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)
	// g.AddNode(n7)
	g.AddUndirectedEdge(n1, n2, 2)
	g.AddUndirectedEdge(n1, n3, 3)
	g.AddUndirectedEdge(n4, n1, 1)
	g.AddUndirectedEdge(n2, n3, 5)
	g.AddUndirectedEdge(n2, n4, 4)
	g.AddUndirectedEdge(n3, n4, 6)
	g.AddUndirectedEdge(n4, n5, 1)
	g.AddUndirectedEdge(n3, n6, 1)
	fmt.Println("...")
	g.Prims()
	fmt.Println("...")
	g.Dijkstra(n2)
	// g.Kruskal()
	// fmt.Println(g.Findk(n1, 5))
	// g.PrintGraph()
	// g.RemoveNode(n3)
	// g.PrintGraph()
	// g.DFS()
	// fmt.Println("No of Paths: ", g.CountAllPaths(n1, n4))
	// g.BFS()
}
