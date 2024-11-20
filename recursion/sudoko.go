package recursion

import "fmt"

func solveSudoko(mat [9][9]int, i, j, n int) bool {
	if i == n {
		// print

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(mat[i][j], " ")
			}
			fmt.Println()
		}
		return true
	}

	//rec case
	if j == n {
		return solveSudoko(mat, i+1, 0, n)
	}

	if mat[i][j] != 0 {
		return solveSudoko(mat, i, j+1, n)
	}

	// cell filling
	// try all possibilities

	for no := 1; no <= n; no++ {
		if isSafe(mat, i, j, no, n) {
			mat[i][j] = no
			solveSubProblem := solveSudoko(mat, i, j+1, n)
			if solveSubProblem {
				return true
			}
		}
	}

	// if no option
	mat[i][j] = 0
	return false
}

func isSafe(mat [9][9]int, i, j, no, n int) bool {
	for k := 0; k < n; k++ {
		if mat[k][j] == no || mat[i][k] == no {
			return false
		}
	}

	sx := (i / 3) * 3
	sy := (j / 3) * 3

	for x := sx; x < sx+3; x++ {
		for y := sy; y < sy+3; y++ {
			if mat[x][y] == no {
				return false
			}
		}
	}
	return true
}

func Sudoko() {
	n := 9
	sudoko := [9][9]int{{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9}}

	if !solveSudoko(sudoko, 0, 0, n) {
		fmt.Println("No Solution Exists")
	}
}
