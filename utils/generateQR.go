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

	writeBits := func(value int, length int) {
		for b := length - 1; b >= 0; b-- {
			result[i] = (value>>b)&1 != 0
			i++
		}
	}

	switch t {
	case "num":
		for pos := 0; pos < len(data); {
			if pos+3 <= len(data) {
				v := int(data[pos]-'0')*100 + int(data[pos+1]-'0')*10 + int(data[pos+2]-'0')
				writeBits(v, 10)
				pos += 3
			} else if pos+2 <= len(data) {
				v := int(data[pos]-'0')*10 + int(data[pos+1]-'0')
				writeBits(v, 7)
				pos += 2
			} else {
				v := int(data[pos] - '0')
				writeBits(v, 4)
				pos += 1
			}
		}
	case "alphanum":
		alphaMap := map[byte]int{
			'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
			'5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
			'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
			'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
			'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35,
			' ': 36, '$': 37, '%': 38, '*': 39, '+': 40, '-': 41, '.': 42, '/': 43, ':': 44,
		}
		for pos := 0; pos < len(data); {
			if pos+2 <= len(data) {
				v := alphaMap[data[pos]]*45 + alphaMap[data[pos+1]]
				writeBits(v, 11)
				pos += 2
			} else {
				v := alphaMap[data[pos]]
				writeBits(v, 6)
				pos += 1
			}
		}
	default:
		for j := range len(data) {
			value := data[j]
			writeBits(int(value), 8)
		}
	}

	for i%8 != 0 {
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
