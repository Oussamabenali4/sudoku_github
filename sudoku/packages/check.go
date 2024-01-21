package sudoku

func Check_Line(line [9]int , number int , column int ) bool {
	// check if the number is unique on the line 
	for i := column + 1 ; i < 9 ; i++ {
		if line[i] == number  {
			return false
		}
	}
	return true
}

func Check_Colone(column [9]int , number int , row int) bool {
	// check if the number is unique in the colon 
	for i := row + 1 ; i < 9 ; i++ {
		if column[i] == number {
			return false
		}
	}
	return true
}
func ScanColumn(sudoku [9][9]int, column int) [9]int{
	columns := [9]int{}
	for i := 0; i < 9; i++ {
		columns[i] = sudoku[i][column]
	}
	return columns
}
func Convert2DemArrayTo1DemArray(array [9][9]int,) {
	line :=[9]int{}
	index :=0
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			line[index] = array[row][col]
		}
	}
}
func Block_Sudoku(Square [9][9]int ,startRow int , startColumn int, Number int) bool {
	count := 0
	for row := startRow; row < 3 + startRow; row++ {
		for col := startColumn; col < 3 + startColumn; col++ {
			if Square[row][col] == Number {
				count++
			}
			if count > 1 {
				return false
			}
		}
	}
	return true
}
func IsValid(sudoku [9][9]int) bool {
	isValid := true
	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			Number := sudoku[row][column]
			if Number != 0 {
				isValid = (Check_Line(sudoku[row],Number,column) && Check_Colone(ScanColumn(sudoku,column),Number,row))
				if isValid == false {
					return false
				}
			}
		}
	}
	return isValid
}

