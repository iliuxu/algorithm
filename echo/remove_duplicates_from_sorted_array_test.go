package echo

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	if 2 != RemoveDuplicates(nums) {
		t.Error("RemoveDuplicates Error!")
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{1, 1, 3, 4, 5, 7}
	for i := 0; i < b.N; i++ {
		RemoveDuplicates(nums)
	}
}
