package echo

import "testing"

func TestSearchInsert(t *testing.T) {
	num := []int{1, 2, 3, 5}
	target := 3
	if 2 != SearchInsert(num, target) {
		t.Error("Search insert error")
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	num := []int{1, 2, 3, 5}
	target := 3
	for i := 0; i < b.N; i++ {
		SearchInsert(num, target)
	}
}
