package p1_test

import (
	"slices"
	"testing"

	"main.go/p1"
)

type DataSet struct {
	nums   []int
	target int
	output []int
}

func TestP1(t *testing.T) {
	testSet := []DataSet{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, val := range testSet {
		tmp := p1.Run(val.nums, val.target)
		if slices.Equal(tmp, val.output) {
			t.Error("Test case failed at slice: ", val.nums)
		}
	}
}
