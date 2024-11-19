package recursion

import "fmt"

func nqueen() {
	boardSize := 4
	board := createMatrix(boardSize, boardSize)
	solveQueen(boardSize, board, 0)
}

func printBoard(board [][]int) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func solveQueen(n int, board [][]int, i int) bool {
	if i == n {
		printBoard(board)
		return true
	}

	// try to place queen in every row
	for j := 0; j < n; j++ {
		// check whether the position is safe or not
		if canPlace(board, n, i, j) {
			board[i][j] = 1
			success := solveQueen(n, board, i+1)

			if success {
				return true
			}

			board[i][j] = 0
		}

	}
	return false
}

func canPlace(board [][]int, n, x, y int) bool {
	//column
	for k := 0; k < x; k++ {
		if board[k][y] == 1 {
			return false
		}
	}

	// left diag

	i := x
	j := y

	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	// right diag

	i = x
	j = y
	for i >= 0 && j < n {
		if board[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	return true
}

func createMatrix(rows, cloumns int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cloumns)
	}

	return matrix
}
