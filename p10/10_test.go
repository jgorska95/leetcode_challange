package p10_test

import (
	"testing"

	"main.go/p10"
)

/*
Example 1:
Input: s = "aa", p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".

Example 2:
Input: s = "aa", p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:
Input: s = "ab", p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
*/

type DataSet struct {
	i1, i2 string
	output bool
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
	}

	for testKey, testVal := range testSet {
		tmp := p10.Run(testVal.i1, testVal.i2)
		if tmp != testVal.output {
			t.Error("Test case failed at test nr: ", testKey)
			t.Errorf("Expected: %t got %t", testVal.output, tmp)
			break
		}
	}
}
