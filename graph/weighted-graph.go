package graph

import "fmt"

type Edge struct {
	To     int
	Weight int
}

type WeightedGraph struct {
	adjList map[int][]Edge
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		adjList: make(map[int][]Edge),
	}
}

func (w *WeightedGraph) AddEdge(from, to, weight int) {
	w.adjList[from] = append(w.adjList[from], Edge{To: to, Weight: weight})
	w.adjList[to] = append(w.adjList[to], Edge{To: from, Weight: weight})
}

func (w *WeightedGraph) GetNeighbours(node int) []Edge {
	return w.adjList[node]
}

func (w *WeightedGraph) Print() {
	for node, edges := range w.adjList {
		fmt.Printf("Node %d: ", node)
		for _, edge := range edges {
			fmt.Printf("-> %d (Weight: %d) ", edge.To, edge.Weight)
		}
		fmt.Println()
	}
}
func weightedGraphImplementation() {
	graph := NewWeightedGraph()
	graph.AddEdge(1, 2, 4)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(2, 3, 2)
	graph.AddEdge(2, 4, 5)
	graph.AddEdge(3, 4, 3)

	graph.Print()
	neighbors := graph.GetNeighbours(2)
	fmt.Println("Neighbors of node 2:")
	for _, neighbor := range neighbors {
		fmt.Printf("-> Node %d with weight %d\n", neighbor.To, neighbor.Weight)
	}

}
