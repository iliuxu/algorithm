package echo

import (
	"testing"
)

var (
	strs = []string{"flower", "flight", "flow"}
	s    = []string{"dog", "racecar", "car"}
)

func TestLongestCommonPrefix(t *testing.T) {
	if "fl" != LongestCommonPrefix(strs) {
		t.Error("longert common prefix error")
	}
	if "" != LongestCommonPrefix(s) {
		t.Error("longert common prefix error")
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix(strs)
	}
}
