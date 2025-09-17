package p54_test

import (
	"testing"

	"main.go/p54"
)

/*
Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

type DataSet struct {
	inputMatrix [][]int
	output      []int
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for testKey, testVal := range testSet {
		tmp := p54.Run(testVal.inputMatrix)
		if len(tmp) != len(testVal.output) {
			t.Error("Test case failed at test nr: ", testKey)
			break
		}
		for key, val := range tmp {
			if val != testVal.output[key] {
				t.Error("Test case failed at test nr: ", testKey)
				break
			}
		}
	}
}
