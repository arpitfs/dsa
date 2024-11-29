package graph

import "fmt"

func dfs() {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(4, 5)
	graph.AddEdge(0, 4)
	graph.AddEdge(3, 4)
	graph.depthFirstSeach(1)
}

func (g *Graph) depthFirstSeachHelper(source int, visited *[]bool) {
	(*visited)[source] = true
	fmt.Print(source, " ,")
	for _, neighbour := range g.adjList[source] {
		if !(*visited)[neighbour] {
			g.depthFirstSeachHelper(neighbour, visited)
		}
	}
}

func (g *Graph) depthFirstSeach(source int) {
	visited := make([]bool, 8)
	g.depthFirstSeachHelper(source, &visited)
}
