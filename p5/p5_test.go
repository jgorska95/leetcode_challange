package p5_test

import (
	"testing"

	"main.go/p5"
)

/*
Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/

type DataSet struct {
	input  string
	output map[string]int
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{"babad", map[string]int{"bab": 1, "aba": 1}},
		{"cbbd", map[string]int{"bb": 1}},
	}

	for key, test := range testSet {
		tmp := p5.Run(test.input)
		if _, ok := test.output[tmp]; !ok {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
