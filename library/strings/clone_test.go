package strings

import (
	stdStrings "strings"
	"testing"
	"unsafe"
)

var emptyString string

func TestClone(t *testing.T) {
	cloneTests := []string{
		"",
		Clone(""),
		stdStrings.Repeat("a", 42)[:0],
		"short",
		stdStrings.Repeat("a", 42),
	}
	for _, input := range cloneTests {
		clone := Clone(input)
		if clone != input {
			t.Errorf("Clone(%q) = %q; want %q", input, clone, input)
		}

		if len(input) != 0 && unsafe.StringData(clone) == unsafe.StringData(input) {
			t.Errorf("Clone(%q) return value should not reference inputs backing memory.", input)
		}

		if len(input) == 0 && unsafe.StringData(clone) != unsafe.StringData(emptyString) {
			t.Errorf("Clone(%#v) return value should be equal to empty string.", unsafe.StringData(input))
		}
	}
}

var stringSink string

func BenchmarkStdClone(b *testing.B) {
	str := stdStrings.Repeat("a", 42)
	b.ReportAllocs()
	for b.Loop() {
		stringSink = stdStrings.Clone(str)
	}
}

func BenchmarkClone(b *testing.B) {
	str := stdStrings.Repeat("a", 42)
	b.ReportAllocs()
	for b.Loop() {
		stringSink = stdStrings.Clone(str)
	}
}
