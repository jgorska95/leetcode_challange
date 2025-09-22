package p9_test

import (
	"testing"

	"main.go/p9"
)

/*
Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

type DataSet struct {
	input  int
	output bool
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{121, true},
		{-121, false},
		{10, false},
	}

	for testKey, testVal := range testSet {
		tmp := p9.Run(testVal.input)
		if tmp != testVal.output {
			t.Error("Test case failed at test nr: ", testKey)
			t.Errorf("Expected: %t got %t", testVal.output, tmp)
			break
		}
	}
}
