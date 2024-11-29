package graph

import "fmt"

func cyclicCheck() {
	graph := NewUnDirectedGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(1, 2)
	// graph.AddEdge(2, 0)
	graph.PrintUndirectedGraph()
	fmt.Println("Is Cycle: ", graph.checkCycle(0))
}

func (g *UnDirectedGraph) PrintUndirectedGraph() {
	for node, neighbors := range g.adjList {
		fmt.Printf("%d -> ", node)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func (g *UnDirectedGraph) checkCycleHelper(parent, source int, visited *[]bool) bool {
	(*visited)[source] = true
	for _, neighbour := range g.adjList[source] {
		if !(*visited)[neighbour] {
			cycleFound := g.checkCycleHelper(parent, neighbour, visited)
			if cycleFound {
				return true
			}
		} else if neighbour != parent {
			return true
		}
	}

	return false
}

func (g *UnDirectedGraph) checkCycle(source int) bool {
	visited := make([]bool, 3)
	return g.checkCycleHelper(-1, source, &visited)
}
