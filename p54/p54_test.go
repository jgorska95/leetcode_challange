package p54_test

import (
	"reflect"
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

	for key, test := range testSet {
		tmp := p54.Run(test.inputMatrix)
		if reflect.DeepEqual(tmp, test.output) {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
