package graph

import "fmt"

func backtrackPath() {

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
	graph.bt(1, 6)

}

func (g *Graph) bt(source int, backtrack int) {
	visited := make([]bool, 8)
	distance := make([]int, 8)
	parent := make([]int, 8)

	for i := 0; i < 8; i++ {
		parent[i] = -1
	}
	queue := []int{source}
	distance[source] = 0
	parent[source] = source
	visited[source] = true
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, neighours := range g.adjList[front] {
			if !visited[neighours] {
				queue = append(queue, neighours)
				parent[neighours] = front
				distance[neighours] = distance[front] + 1
				visited[neighours] = true
			}
		}
	}
	temp := backtrack
	for temp != source {
		fmt.Print(temp, "--")
		temp = parent[temp]
	}
	fmt.Print(temp)
	fmt.Println()
}
