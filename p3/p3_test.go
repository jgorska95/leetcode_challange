package p3_test

import (
	"testing"

	"main.go/p3"
)

type DataSet struct {
	pharse      string
	expectedLen int
}

func TestP3(t *testing.T) {
	testSet := []DataSet{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
	}

	for key, test := range testSet {
		tmp := p3.Run(test.pharse)
		if tmp != test.expectedLen {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
