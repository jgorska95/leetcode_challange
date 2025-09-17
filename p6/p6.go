package p6

func Run(s string, numRows int) string {
	return convert(s, numRows)
}

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) == 0 {
		return s
	}
	retByteSlice := []byte{}

	verticalCount := numRows
	diagonalCount := numRows - 2

	offset := verticalCount + diagonalCount

	for row := 0; row < numRows; row++ {
		id := row

		vertical := true

		diagonalOffset := offset - row*2
		verticalOffset := offset - diagonalOffset

		for id < len(s) {
			retByteSlice = append(retByteSlice, s[id])

			if row == 0 || row == numRows-1 {
				id += offset
			} else {
				if vertical {
					id += diagonalOffset
				} else {
					id += verticalOffset
				}
			}

			vertical = !vertical
		}
	}

	return string(retByteSlice)
}
