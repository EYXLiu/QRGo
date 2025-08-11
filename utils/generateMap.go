package utils

func GenerateMap(matrix [][]bool, data []bool) [][]bool {
	col, row := len(matrix)-1, len(matrix)-1
	upward := true

	next := func() {
		if upward {
			if col%2 == 0 {
				col -= 1
			} else {
				col += 1
				row -= 1
			}
		} else {
			if col%2 == 0 {
				col -= 1
			} else {
				col += 1
				row += 1
			}
		}
		if (row == 8 && col > 21) || (row == -1) {
			upward = !upward
			row += 1
			col -= 2
		}
		if row > 28 {
			upward = !upward
			row -= 1
			col -= 2
		}
		if row == 24 && col == 24 {
			row -= 5
		}
		if row == 20 && col == 22 {
			row += 5
		}
		if col == 20 && row <= 24 && row >= 20 {
			col -= 1
		}
		if row == 6 {
			if upward {
				row -= 1
			} else {
				row += 1
			}
		}
	}
	for i := range data {
		matrix[row][col] = data[i]
		next()
	}

	return matrix
}
