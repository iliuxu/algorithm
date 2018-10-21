package echo

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{1,2,3}
	val := 3
	if 2 != RemoveElement(nums,val) {
		t.Error("Remove element error")
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{1,2,3}
	val := 3
	for i:=0; i<b.N; i++ {
		RemoveElement(nums, val)
	}
}
