package p7_test

import (
	"testing"

	"main.go/p7"
)

/*
Example 1:
Input: x = 123
Output: 321

Example 2:
Input: x = -123
Output: -321

Example 3:
Input: x = 120
Output: 21
*/

type DataSet struct {
	input  int
	output int
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for testKey, testVal := range testSet {
		tmp := p7.Run(testVal.input)
		if tmp != testVal.output {
			t.Error("Test case failed at test nr: ", testKey)
			break
		}
	}
}
