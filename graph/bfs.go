package graph

func bfs() {

	graph := NewGraphList(4, false)
	graph.AddEdge(0, 1, 0)
	graph.AddEdge(1, 2, 0)
	graph.Display()

}
