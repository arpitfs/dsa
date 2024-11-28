package graph

import (
	"fmt"
	"math"
)

func implementGraph() {
	g := NewGraphList(4, false)
	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 2, 0)
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)

	g.Display()
	g1 := NewGraphMatrix(4, true)
	g1.AddEdge(0, 1, 10)
	g1.AddEdge(0, 2, 20)
	g1.AddEdge(1, 2, 30)
	g1.AddEdge(2, 3, 40)

	g1.Display()

}

type GraphList struct {
	V        int
	Adj      map[int][]int
	weighted bool
}

func NewGraphList(v int, weighted bool) *GraphList {
	return &GraphList{
		V:        v,
		Adj:      make(map[int][]int),
		weighted: weighted,
	}
}

func (g *GraphList) AddEdge(u, v, weight int) {
	if g.weighted {
		g.Adj[u] = append(g.Adj[u], v)
		g.Adj[v] = append(g.Adj[v], u)
	} else {
		g.Adj[u] = append(g.Adj[u], v)
		g.Adj[v] = append(g.Adj[v], u) //
	}
}

func (g *GraphList) Display() {
	for i := 0; i < g.V; i++ {
		fmt.Printf("Vertex %d: ", i)
		for _, v := range g.Adj[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func (g *GraphList) AddEdgeDirected(u, v int, weight int) {
	if g.weighted {
		g.Adj[u] = append(g.Adj[u], v)
	} else {
		g.Adj[u] = append(g.Adj[u], v)
	}
}

type GraphMatrix struct {
	V        int
	Adj      [][]int
	Weighted bool
}

func NewGraphMatrix(v int, weighted bool) *GraphMatrix {
	adjMatrix := make([][]int, v)
	for i := 0; i < v; i++ {
		adjMatrix[i] = make([]int, v)
		if weighted {
			for j := 0; j < v; j++ {
				adjMatrix[i][j] = math.MaxInt32
			}
		}
	}
	return &GraphMatrix{V: v, Adj: adjMatrix, Weighted: weighted}
}

func (g *GraphMatrix) AddEdge(u, v, weight int) {
	if g.Weighted {
		g.Adj[u][v] = weight
		g.Adj[v][u] = weight
	} else {
		g.Adj[u][v] = 1
		g.Adj[v][u] = 1
	}
}

func (g *GraphMatrix) Display() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < g.V; j++ {
			if g.Adj[i][j] != math.MaxInt32 {
				fmt.Printf("%d ", g.Adj[i][j])
			} else {
				fmt.Print("∞ ")
			}
		}
		fmt.Println()
	}
}
