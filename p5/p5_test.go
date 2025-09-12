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
	output string
}

func TestP4(t *testing.T) {
	testSet := []DataSet{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for key, test := range testSet {
		tmp := p5.Run(test.input)
		if tmp != test.output {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
