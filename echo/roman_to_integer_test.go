package echo

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	if 1994 != RomanToInt("MCMXCIV") {
		t.Error("1994 != RomanToInt(\"MCMXCIV\")")
	}
	if 58 != RomanToInt("LVIII") {
		t.Error("58 != RomanToInt(\"LVIII\")")
	}

	if 9 != RomanToInt("IX") {
		t.Error("9 != RomanToInt(\"IX\")")
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("M")
	}
}
