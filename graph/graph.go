package graph

import (
	"fmt"
	"github.com/shubhamagarwal003/go-algos/queue"
	"reflect"
)

type Edge struct {
	dest *Node
}

type Node struct {
	value int
	edges []*Edge
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

func (g *Graph) AddUndirectedEdge(node1 *Node, node2 *Node) {
	e1 := &Edge{dest: node1}
	e2 := &Edge{dest: node2}
	node1.edges = append(node1.edges, e2)
	node2.edges = append(node2.edges, e1)
}

func (g *Graph) AddDirectedEdge(node1 *Node, node2 *Node) {
	e := &Edge{dest: node2}
	node1.edges = append(node1.edges, e)
}

func (g *Graph) RemoveNode(node *Node) {
	i := 0
	for _, n := range g.nodes {
		if n == node {
			break
		}
		i++
	}
	if i < len(g.nodes) {
		g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
		for _, n := range g.nodes {
			j := 0
			for _, e := range n.edges {
				if e.dest == node {
					n.edges = append(n.edges[:j], n.edges[j+1:]...)
					j--
				}
				j++
			}
		}
	}
}

func (g *Graph) PrintGraph() {
	for _, node := range g.nodes {
		fmt.Print(node.value)
		fmt.Print("->")
		for _, edge := range node.edges {
			fmt.Print(" ")
			fmt.Print(edge.dest.value)
		}
		fmt.Println("")
	}
}

func (g *Graph) DFS() {
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	for _, node := range g.nodes {
		if !visited[node] {
			g.DFSUtil(node, visited)
		}
	}
}

func (g *Graph) DFSUtil(node *Node, visited map[*Node]bool) {
	fmt.Println(node.value)
	visited[node] = true
	for _, e := range node.edges {
		if !visited[e.dest] {
			fmt.Println(e.dest.value)
			visited[e.dest] = true
			g.DFSUtil(e.dest, visited)
		}
	}
}

func (g *Graph) BFS() {
	q := queue.NewQueue()
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	for _, node := range g.nodes {
		g.BFSUtil(node, q, visited)
	}
}

func (g *Graph) BFSUtil(node *Node, q *queue.Queue, visited map[*Node]bool) {
	if !visited[node] {
		q.Push(node)
		visited[node] = true
	}
	n, _ := q.Pop()
	for n != nil {
		nVal := reflect.ValueOf(n)
		nod, _ := nVal.Interface().(*Node)
		fmt.Println(nod.value)
		for _, e := range nod.edges {
			if !visited[e.dest] {
				q.Push(e.dest)
				visited[e.dest] = true
			}
		}
		n, _ = q.Pop()
	}
}

func NewGraph() *Graph {
	g := &Graph{}
	return g
}

func NewNode(val int) *Node {
	node := &Node{value: val}
	return node
}
