package graph

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) PrintGraph() {
	for node, neighbors := range g.adjList {
		fmt.Printf("%d -> ", node)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func bfs() {

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
	graph.breadthFirstSeach(1)

}

func (g *Graph) breadthFirstSeach(source int) {
	queue := []int{source}
	visited := make([]bool, 8)
	visited[source] = true
	fmt.Println("Breadth First Seach")
	for len(queue) > 0 {
		front := queue[0]
		fmt.Print(front, " ,")
		queue = queue[1:]
		for _, neighours := range g.adjList[front] {
			if !visited[neighours] {
				queue = append(queue, neighours)
				visited[neighours] = true
			}
		}
	}
}
