package p6_test

import (
	"testing"

	"main.go/p6"
)

/*
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"
*/

type DataSet struct {
	inputString string
	inputInt    int
	output      string
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}

	for key, test := range testSet {
		tmp := p6.Run(test.inputString, test.inputInt)
		if tmp != test.output {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
