package p4_test

import (
	"testing"

	"main.go/p4"
)

/*
Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/

type DataSet struct {
	nums1  []int
	nums2  []int
	Output float64
}

func Test(t *testing.T) {
	testSet := []DataSet{
		{[]int{1, 3}, []int{2}, 2.00000},
		{[]int{1, 2}, []int{3, 4}, 2.50000},
	}

	for key, test := range testSet {
		tmp := p4.Run(test.nums1, test.nums2)
		if tmp != test.Output {
			t.Error("Test case failed at test nr: ", key)
		}
	}
}
