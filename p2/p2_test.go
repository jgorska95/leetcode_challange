package p2_test

import (
	"testing"

	"main.go/p2"
)

/*
Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/

type DataSet struct {
	l1     *p2.ListNode
	l2     *p2.ListNode
	Output *p2.ListNode
}

func Equal(c1, c2 *p2.ListNode) bool {
	if c1.Val != c2.Val {
		return false
	}
	if c1.Val == c2.Val && c1.Next != nil && c2.Next != nil {
		return Equal(c1.Next, c2.Next)
	}
	if c1.Next != nil && c2.Next == nil {
		return false
	}
	if c1.Next == nil && c2.Next != nil {
		return false
	}
	return true
}

func TestP2(t *testing.T) {
	testSet := []DataSet{
		{&p2.ListNode{2, &p2.ListNode{4, &p2.ListNode{3, nil}}}, &p2.ListNode{5, &p2.ListNode{6, &p2.ListNode{4, nil}}}, &p2.ListNode{7, &p2.ListNode{0, &p2.ListNode{8, nil}}}},
		{&p2.ListNode{0, nil}, &p2.ListNode{0, nil}, &p2.ListNode{0, nil}},
		{&p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, nil}}}}}}}, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, nil}}}}, &p2.ListNode{8, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{9, &p2.ListNode{0, &p2.ListNode{0, &p2.ListNode{0, &p2.ListNode{1, nil}}}}}}}}},
	}

	for _, val := range testSet {
		tmp := p2.Run(val.l1, val.l2)
		if !Equal(tmp, val.Output) {
			t.Error("Test case failed at: ", val.Output)
		}
	}
}
