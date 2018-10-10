package echo

import (
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {

}

func BenchmarkMeregeTwoSortedLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			l1 = &ListNode{1, nil}
			l2 = &ListNode{2, nil}
		)
		MergeTwoSortedLists(l1, l2)
	}
}
