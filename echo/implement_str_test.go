package echo

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	if 0 != StrStr("s", "s") {
		t.Error(" 0 != StrStr(\"s\", \"s\")")
	}
	if -1 != StrStr("aaaaa", "bba") {
		t.Error("StrStr Error")
	}
}

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrStr("aaaaa", "bba")
	}
}
