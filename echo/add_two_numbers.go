package echo

type ListNode struct {
	Val  int
	Next *ListNode
}

// Add Two Numbers
//
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
//
//  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//  Output: 7 -> 0 -> 8
//  Explanation: 342 + 465 = 807.
// https://leetcode.com/problems/add-two-numbers/description/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	p1 := l1
	p2 := l2
	newHead := &ListNode{0, nil}
	p3 := newHead
	for p1 != nil || p2 != nil {
		if p1 != nil {
			carry += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			carry += p2.Val
			p2 = p2.Next
		}
		p3.Next = &ListNode{carry % 10, nil}
		p3 = p3.Next
		carry /= 10
	}
	if carry == 1 {
		p3.Next = &ListNode{1, nil}
	}
	return newHead.Next
}
