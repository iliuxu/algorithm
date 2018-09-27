package echo

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	if 321 != ReverseInteger(123) {
		t.Error("321 != ReverseInteger(123)")
	}
	if -321 != ReverseInteger(-123) {
		t.Error("-321 != ReverseInteger(-123)")
	}
	if 21 != ReverseInteger(120) {
		t.Error("21 != ReverseInteger(120)")
	}
}

func BenchmarkReverseInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseInteger(123)
	}
}
