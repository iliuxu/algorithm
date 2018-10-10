package echo

//Merge Two Sorted Lists
//
// Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
// Example:
//
//  Input: 1->2->4, 1->3->4
//  Output: 1->1->2->3->4->4
// https://leetcode.com/problems/merge-two-sorted-lists/description/
func MergeTwoSortedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	p3 := ListNode{0, nil}
	newHead := &p3
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			newHead.Next = &ListNode{p1.Val, nil}
			p1 = p1.Next
		} else {
			newHead.Next = &ListNode{p2.Val, nil}
			p2 = p2.Next
		}
		newHead = newHead.Next
	}
	if p1 != nil {
		newHead.Next = p1
	}
	if p2 != nil {
		newHead.Next = p2
	}
	return p3.Next
}
