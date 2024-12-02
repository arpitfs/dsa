package graph

// import "fmt"

// var directions = [][2]int{
// 	{0, 1},  // right
// 	{1, 0},  // down
// 	{0, -1}, // left
// 	{-1, 0}, // up
// }

// func largestIsland() {
// 	grid := [][]int{
// 		{1, 1, 0, 0, 0},
// 		{1, 1, 0, 0, 1},
// 		{0, 0, 0, 1, 1},
// 		{0, 0, 0, 0, 0},
// 		{1, 1, 1, 0, 0},
// 	}

// 	largest := 0

// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[0]); j++ {

// 			if grid[i][j] == 1 {
// 				islandSize := dfsIsland(grid, i, j)
// 				if islandSize > largest {
// 					largest = islandSize
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println("Island:", largest)
// }

// func dfsIsland(grid [][]int, x, y int) int {
// 	if x < 0 || x >= len(grid) || y < 0 || y > len(grid) || grid[x][y] == 0 {
// 		return 0
// 	}

// 	grid[x][y] = 0
// 	size := 1
// 	for _, dir := range directions {
// 		newX, newY := x+dir[0], y+dir[1]
// 		size += dfsIsland(grid, newX, newY)
// 	}

// 	return size
// }
