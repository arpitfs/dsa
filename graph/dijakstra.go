package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// min-heap
type PQ []*Node

type Node struct {
	Id       int
	Distance int
}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (w *WeightedGraph) dijkstra(start int) map[int]int {
	distance := make(map[int]int)
	for node := range w.adjList {
		distance[node] = math.MaxInt
	}

	distance[start] = 0

	pq := make(PQ, 0)
	heap.Push(&pq, &Node{Id: start, Distance: 0})

	visited := make(map[int]bool)
	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)
		if visited[node.Id] {
			continue
		}
		visited[node.Id] = true

		for _, edge := range w.adjList[node.Id] {
			if !visited[edge.To] {
				newDist := node.Distance + edge.Weight
				if newDist < distance[edge.To] {
					distance[edge.To] = newDist
					heap.Push(&pq, &Node{Id: edge.To, Distance: newDist})
				}
			}
		}
	}

	return distance

}

func dijkstraAlgo() {
	graph := NewWeightedGraph()
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(2, 3, 2)
	graph.AddEdge(2, 4, 5)
	graph.AddEdge(3, 4, 1)

	startNode := 1
	distances := graph.dijkstra(startNode)

	for node, dist := range distances {
		fmt.Printf("Node %d: %d\n", node, dist)
	}
}
