package sudoku

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

// func Solve_Soduko(array [9][9]int, x int, y int) [9][9]int {
// 	for j := 0; j < 9; j++ {
// 		for i := 0; i < 9; i++ {
// 			if array[i][j] == 0 {
// 				for solution := 1; solution <= 9; solution++ {
// 					array[i][j] = solution
// 					if IsValid(array) {
// 						Stro_Changes(i, j, solution)
// 						break
// 					} else {
// 						continue
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return array
// }

// ////////////////////////////////////////////////
// type Coordinate struct {
// 	Row, Col int
// }

// func Stro_Changes(i int, j int, solution int) map[Coordinate]int {

//     changes := make(map[Coordinate]int)
//     changes[Coordinate{i, j}] = solution

//     return changes
// }
// //////////////////////////////////////////////

// func Back_Track(){

// }
