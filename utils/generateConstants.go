package utils

func GenerateDefaults(matrix [][]bool) [][]bool {
	//Position Squares
	for i := range 7 {
		matrix[0][i] = true
		matrix[6][i] = true
		matrix[i][0] = true
		matrix[i][6] = true

		matrix[0][22+i] = true
		matrix[6][22+i] = true
		matrix[22+i][0] = true
		matrix[22+i][6] = true

		matrix[22][i] = true
		matrix[28][i] = true
		matrix[i][22] = true
		matrix[i][28] = true
	}
	for i := 2; i <= 4; i++ {
		for j := 2; j <= 4; j++ {
			matrix[i][j] = true
			matrix[i][22+j] = true
			matrix[22+i][j] = true
		}
	}

	//Alignment Pattern
	for i := range 5 {
		matrix[20][20+i] = true
		matrix[24][20+i] = true
		matrix[20+i][20] = true
		matrix[20+i][24] = true
	}
	matrix[22][22] = true

	//Timing Strips
	for i := range 7 {
		matrix[6][8+i*2] = true
		matrix[8+i*2][6] = true
	}

	//Pixel Always Dark
	matrix[21][8] = true

	//Error Correction Level
	coordsMap := map[string][][2]int{
		"TopLeft":    {{8, 0}, {8, 1}, {8, 2}, {8, 4}, {8, 5}, {8, 7}, {8, 8}, {7, 8}, {2, 8}},
		"TopRight":   {{8, 21}, {8, 22}, {8, 26}},
		"BottomLeft": {{21, 8}, {22, 8}, {23, 8}, {24, 8}, {26, 8}, {27, 8}, {28, 8}},
	}

	for _, coords := range coordsMap {
		for _, coord := range coords {
			matrix[coord[0]][coord[1]] = true
		}
	}

	return matrix
}

func GenerateQuietZone(matrix [][]bool) [][]bool {
	//Surround the QR Code with white
	size := len(matrix)
	newSize := size + 2
	newMatrix := make([][]bool, newSize)
	for i := range newMatrix {
		newMatrix[i] = make([]bool, newSize)
	}

	for i := range size {
		for j := range size {
			newMatrix[i+1][j+1] = matrix[i][j]
		}
	}

	return newMatrix
}
