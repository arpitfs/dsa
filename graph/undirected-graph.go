package graph

import "fmt"

type UnDirectedGraph struct {
	adjList map[int][]int
}

func NewUnDirectedGraph() *UnDirectedGraph {
	return &UnDirectedGraph{adjList: make(map[int][]int)}
}

func (g *UnDirectedGraph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *UnDirectedGraph) PrintGraph() {
	for node, neighbors := range g.adjList {
		fmt.Printf("%d -> ", node)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func unDirectedGraph() {

	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(4, 5)
	graph.AddEdge(0, 4)
	graph.AddEdge(3, 4)
	graph.PrintGraph()

}
