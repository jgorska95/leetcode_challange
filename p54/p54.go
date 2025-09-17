package p54

func Run(matrix [][]int) []int {
	return spiralOrder(matrix)
}

func spiralOrder(matrix [][]int) []int {
	retSlice := []int{}

	rowMax := len(matrix)
	colMax := len(matrix[0])

	row, col := 0, 0

	totalCnt := rowMax * colMax
	cnt := 1

	for cnt <= totalCnt {
		for i := col; i < colMax && cnt <= totalCnt; i++ {
			cnt++
			retSlice = append(retSlice, matrix[row][i])
		}
		row++

		for i := row; i < rowMax && cnt <= totalCnt; i++ {
			cnt++
			retSlice = append(retSlice, matrix[i][colMax-1])
		}
		colMax--

		for i := colMax - 1; i >= col && cnt <= totalCnt; i-- {
			cnt++
			retSlice = append(retSlice, matrix[rowMax-1][i])
		}
		rowMax--

		for i := rowMax - 1; i >= row && cnt <= totalCnt; i-- {
			cnt++
			retSlice = append(retSlice, matrix[i][col])
		}
		col++
	}

	return retSlice
}
