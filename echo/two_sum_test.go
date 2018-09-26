package echo

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3}
	var answer []int
	answer = []int{1, 2}
	if answer[0] != TwoSum(nums, 5)[0] && answer[1] != TwoSum(nums, 5)[1] {
		t.Error("error")
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		TwoSum(nums, 5)
	}

}
