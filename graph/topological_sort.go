package graph

import (
	"fmt"
	"github.com/shubhamagarwal003/go-algos/queue"
	"github.com/shubhamagarwal003/go-algos/stack"
	"reflect"
)

func (g *Graph) TopologicalSort() {
	st := stack.NewStack()
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	for _, node := range g.nodes {
		if !visited[node] {
			g.topologicalSortUtil(node, visited, st)
		}
	}
	node, _ := st.Pop()
	for node != nil {
		nodeVal, _ := reflect.ValueOf(node).Interface().(*Node)
		fmt.Println(nodeVal.value)
		node, _ = st.Pop()
	}
}

func (g *Graph) topologicalSortUtil(node *Node, visited map[*Node]bool, st *stack.Stack) {
	visited[node] = true
	for _, e := range node.edges {
		if !visited[e.dest] {
			g.topologicalSortUtil(e.dest, visited, st)
		}
	}
	st.Push(node)
}

func (g *Graph) TopologicalSortKahn() {
	indegrees := g.findInDegrees()
	visited := make(map[*Node]bool)
	for _, node := range g.nodes {
		visited[node] = false
	}
	q := queue.NewQueue()
	for q.Length() != len(g.nodes) {
		flag := false
		for node, deg := range indegrees {
			if deg == 0 && !visited[node] {
				q.Push(node)
				visited[node] = true
				for _, e := range node.edges {
					indegrees[e.dest]--
					flag = true
				}
			}
		}
		if !flag {
			break
		}
	}
	if q.Length() != len(g.nodes) {
		fmt.Println("Topological Sort not possible")
	} else {
		n, _ := q.Pop()
		for n != nil {
			nVal := reflect.ValueOf(n)
			nod, _ := nVal.Interface().(*Node)
			fmt.Println(nod.value)
			n, _ = q.Pop()
		}
	}
}

func (g *Graph) findInDegrees() map[*Node]int {
	indegrees := make(map[*Node]int)
	for _, node := range g.nodes {
		indegrees[node] = 0
	}

	for _, node := range g.nodes {
		for _, e := range node.edges {
			indegrees[e.dest]++
		}
	}
	return indegrees
}
