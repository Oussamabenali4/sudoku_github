package main

import (
	"fmt"
	"os"

	sudoku "sudoku/packages"
)

func solveSudoku(board *[9][9]int, isValid func([9][9]int) bool) bool {
	row, col := -1, -1
	isEmpty := true

	// Find the first empty cell
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				isEmpty = false
				break
			}
		}
		if !isEmpty {
			break
		}
	}

	// If no empty cell is left, return true if the board is valid
	if isEmpty {
		return isValid(*board)
	}
	/////////////////////////////////////////////////////////////////////////////////////////// here i should start solving the sudoku
	// Try filling the cell with numbers 1 to 9
	for num := 1; num <= 9; num++ {
		board[row][col] = num
		if isValid(*board) && solveSudoku(board, isValid) {
			return true
		}
		board[row][col] = 0 // Backtrack
	}
	return false
}

func main() {
	var line [9][9]int
	arguments := os.Args[1:]
	// )
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			line[i][j] = sudoku.Parser(arguments)[i][j]
		}
	}

	if sudoku.IsValid(line) {
		// Attempt to solve the Sudoku
		if solveSudoku(&line, sudoku.IsValid) {
			// Print the solved Sudoku
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Printf("%d ", line[i][j])
				}
				fmt.Println()
			}
		}
	} else {
		fmt.Println("Error")
	}
}
