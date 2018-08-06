package main

import (
	"fmt"
	// "github.com/shubhamagarwal003/go-algos/graph"
	"github.com/shubhamagarwal003/go-algos/avl"
	// "reflect"
)

func main() {
	a := avl.NewAvlNode(4)
	a = a.Insert(3)
	a = a.Insert(5)
	a = a.Insert(1)
	a = a.Insert(2)
	a.Print()
	// fmt.Println("a", a)
	// fmt.Println("a", a)
	// a.Print()
}

/*func main() {
	t := trie.NewTrieNode()
	t.Insert("there")
	t.Insert("their")
	t.Insert("these")
	t.Insert("the")
	fmt.Println("....")
	fmt.Println(t.Find("there"))
	fmt.Println(t.Find("the"))
	fmt.Println(t.Find("ant"))
	t.Delete("the")
	fmt.Println(t.Find("the"))
	t.Delete("there")
	fmt.Println(t.Find("there"))
	t.Print()
}*/

/*func main() {
	b := heap.NewHeap()
	b.Insert(4)
	b.Insert(1)
	b.Insert(3)
	b.Insert(14)
	b.Insert(2)
	b.Insert(10)
	fmt.Println(b)
	b.DecreaseKey(2, 5)
	fmt.Println(b)
	// a, err := b.ExtractMin()
	// for err == nil {
	// 	fmt.Println(a)
	// 	a, err = b.ExtractMin()
	// }
}*/

/*func main() {
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
}*/
