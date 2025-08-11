package utils

func GenerateQRCode(data string) [][]bool {
	matrix := make([][]bool, 29)
	for i := range matrix {
		matrix[i] = make([]bool, 29)
	}

	t := "byte"
	if isAlphanumeric(data) {
		if isNumeric(data) {
			t = "num"
		} else {
			t = "alphanum"
		}
	}
	matrix = GenerateDefaults(matrix)

	length := 440
	result := make([]bool, length)

	//Get Mode
	switch t {
	case "num":
		result[0] = false
		result[1] = false
		result[2] = false
		result[3] = true
	case "alphanum":
		result[0] = false
		result[1] = false
		result[2] = true
		result[3] = false
	default:
		result[0] = false
		result[1] = true
		result[2] = false
		result[3] = false
	}
	//Get Length
	i := 4
	v := len(data)
	for i < 12 {
		result[11-i+4] = (v & 1) != 0
		v >>= 1
		i++
	}

	switch t {
	case "num":

	case "alphanum":

	default:
		for j := range len(data) {
			value := data[j]
			for w := range 8 {
				_ = w
				result[i+7-w] = (value & 1) != 0
				value >>= 1
			}
			i += 8
		}
	}

	for i % 8 != 0 {
		result[i] = false
		i++
	}

	//Pad rest of the values
	pattern := "1110110000010001"
	patternLen := len(pattern)
	start := 0
	for i < 440 {
		result[i] = pattern[start%patternLen] != '0'
		start++
		i++
	}

	matrix = GenerateMap(matrix, result)
	//matrix = GenerateQuietZone(matrix)

	return matrix
}
