package sudoku

func Parser(arguments []string) [9][9]int {
	// var sl []string
	var line [9][9]int
	for j, arg := range arguments {
		for i, char := range arg {
			if char == '.' {
				line[j][i] = 0
			} else {
				line[j][i] = int(rune(char) - '0')
			}
		}
	}

	return line
}
