package echo

import (
	"testing"
)

var (
	l1 = &ListNode{1, nil}
	l2 = &ListNode{2, nil}
)

func TestAddTwoNumbers(t *testing.T) {
	l3 := &ListNode{3, nil}
	if l3.Val != AddTwoNumbers(l1, l2).Val {
		t.Error("AddTwoNumbers Error")
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(l1, l2)
	}
}
