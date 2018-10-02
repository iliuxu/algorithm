package echo

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	if true != IsValid("{}") {
		t.Error("is valid error")
	}
	if true != IsValid("(){}[]") {
		t.Error("is valid error")
	}
	if false != IsValid("{[][]}}") {
		t.Error("is valid error")
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("{{[]}}")
	}
}
