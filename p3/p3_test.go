package p3

import (
	"testing"
)

type DataSet struct {
	pharse      string
	expectedLen int
}

func TestP3(t *testing.T) {
	testSet := []DataSet{{"abcabcbb", 3}, {"bbbbb", 1}, {"pwwkew", 3}, {"dvdf", 3}}

	for _, val := range testSet {
		tmp := Run(val.pharse)
		if tmp != val.expectedLen {
			t.Error("Test case failed at: ", val.pharse)
		}
	}
}
