package echo

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if true != IsPalindrome(121) {
		t.Error("true != IsPalindrome(121)")
	}
	if true == IsPalindrome(-121) {
		t.Error("true == IsPalindrome(-121)")
	}
	if true != IsPalindrome(0) {
		t.Error("true != IsPalindrome(0)")
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome(i)
	}
}
