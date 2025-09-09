package p2

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func Run(l1 *ListNode, l2 *ListNode) {
	_ = addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1ptr := l1
	l2ptr := l2

	dummy := &ListNode{}
	Output := dummy

	carry := 0

	for l1ptr != nil || l2ptr != nil {
		sum := carry

		if l1ptr != nil {
			sum += l1ptr.Val
			l1ptr = l1ptr.Next
		}

		if l2ptr != nil {
			sum += l2ptr.Val
			l2ptr = l2ptr.Next
		}

		carry = sum / 10
		digit := sum % 10

		Output.Next = &ListNode{Val: digit}
		Output = Output.Next
	}

	if carry > 0 {
		Output.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
