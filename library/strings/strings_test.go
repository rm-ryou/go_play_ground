package strings

import (
	"math/rand"
	stdStrings "strings"
	"testing"
)

const BenchmarkString = "some_text=some☺value"

var ContainsTests = []struct {
	str, substr string
	expected    bool
}{
	{"abc", "bc", true},
	{"abc", "bcd", false},
	{"abc", "", true},
	{"", "a", false},

	// cases to cover code in runtime/asm_amd64.s:indexShortStr
	// 2-byte needle
	{"xxxxxx", "01", false},
	{"01xxxx", "01", true},
	{"xx01xx", "01", true},
	{"xxxx01", "01", true},
	{"01xxxxx"[1:], "01", false},
	{"xxxxx01"[:6], "01", false},
	// 3-byte needle
	{"xxxxxxx", "012", false},
	{"012xxxx", "012", true},
	{"xx012xx", "012", true},
	{"xxxx012", "012", true},
	{"012xxxxx"[1:], "012", false},
	{"xxxxx012"[:7], "012", false},
	// 4-byte needle
	{"xxxxxxxx", "0123", false},
	{"0123xxxx", "0123", true},
	{"xx0123xx", "0123", true},
	{"xxxx0123", "0123", true},
	{"0123xxxxx"[1:], "0123", false},
	{"xxxxx0123"[:8], "0123", false},
	// 5-7-byte needle
	{"xxxxxxxxx", "01234", false},
	{"01234xxxx", "01234", true},
	{"xx01234xx", "01234", true},
	{"xxxx01234", "01234", true},
	{"01234xxxxx"[1:], "01234", false},
	{"xxxxx01234"[:9], "01234", false},
	// 8-byte needle
	{"xxxxxxxxxxxx", "01234567", false},
	{"01234567xxxx", "01234567", true},
	{"xx01234567xx", "01234567", true},
	{"xxxx01234567", "01234567", true},
	{"01234567xxxxx"[1:], "01234567", false},
	{"xxxxx01234567"[:12], "01234567", false},
	// 9-15-byte needle
	{"xxxxxxxxxxxxx", "012345678", false},
	{"012345678xxxx", "012345678", true},
	{"xx012345678xx", "012345678", true},
	{"xxxx012345678", "012345678", true},
	{"012345678xxxxx"[1:], "012345678", false},
	{"xxxxx012345678"[:13], "012345678", false},
	// 16-byte needle
	{"xxxxxxxxxxxxxxxxxxxx", "0123456789ABCDEF", false},
	{"0123456789ABCDEFxxxx", "0123456789ABCDEF", true},
	{"xx0123456789ABCDEFxx", "0123456789ABCDEF", true},
	{"xxxx0123456789ABCDEF", "0123456789ABCDEF", true},
	{"0123456789ABCDEFxxxxx"[1:], "0123456789ABCDEF", false},
	{"xxxxx0123456789ABCDEF"[:20], "0123456789ABCDEF", false},
	// 17-31-byte needle
	{"xxxxxxxxxxxxxxxxxxxxx", "0123456789ABCDEFG", false},
	{"0123456789ABCDEFGxxxx", "0123456789ABCDEFG", true},
	{"xx0123456789ABCDEFGxx", "0123456789ABCDEFG", true},
	{"xxxx0123456789ABCDEFG", "0123456789ABCDEFG", true},
	{"0123456789ABCDEFGxxxxx"[1:], "0123456789ABCDEFG", false},
	{"xxxxx0123456789ABCDEFG"[:21], "0123456789ABCDEFG", false},

	// partial match cases
	{"xx01x", "012", false},                             // 3
	{"xx0123x", "01234", false},                         // 5-7
	{"xx01234567x", "012345678", false},                 // 9-15
	{"xx0123456789ABCDEFx", "0123456789ABCDEFG", false}, // 17-31, issue 15679
}

func TestContains(t *testing.T) {
	for _, ct := range ContainsTests {
		if Contains(ct.str, ct.substr) != ct.expected {
			t.Errorf("Contains(%s, %s) = %v, want %v",
				ct.str, ct.substr, !ct.expected, ct.expected)
		}
	}
}

func BenchmarkStdIndex(b *testing.B) {
	if got := stdStrings.Index(BenchmarkString, "v"); got != 17 {
		b.Fatalf("wrong index: expected 17, got=%d", got)
	}
	for b.Loop() {
		stdStrings.Index(BenchmarkString, "v")
	}
}

func BenchmarkIndex(b *testing.B) {
	if got := Index(BenchmarkString, "v"); got != 17 {
		b.Fatalf("wrong index: expected 17, got=%d", got)
	}
	for b.Loop() {
		Index(BenchmarkString, "v")
	}
}

func makeBenchInputHard() string {
	tokens := [...]string{
		"<a>", "<p>", "<b>", "<strong>",
		"</a>", "</p>", "</b>", "</strong>",
		"hello", "world",
	}
	x := make([]byte, 0, 1<<20)
	r := rand.New(rand.NewSource(99))
	for {
		i := r.Intn(len(tokens))
		if len(x)+len(tokens[i]) >= 1<<20 {
			break
		}
		x = append(x, tokens[i]...)
	}
	return string(x)
}

var benchInputHard = makeBenchInputHard()

func benchmarkStdIndexHard(b *testing.B, sep string) {
	for b.Loop() {
		stdStrings.Index(benchInputHard, sep)
	}
}

func benchmarkIndexHard(b *testing.B, sep string) {
	for b.Loop() {
		Index(benchInputHard, sep)
	}
}

func BenchmarkStdIndexHard1(b *testing.B) { benchmarkStdIndexHard(b, "<>") }
func BenchmarkIndexHard1(b *testing.B)    { benchmarkIndexHard(b, "<>") }
func BenchmarkStdIndexHard2(b *testing.B) { benchmarkStdIndexHard(b, "</pre>") }
func BenchmarkIndexHard2(b *testing.B)    { benchmarkIndexHard(b, "</pre>") }
func BenchmarkStdIndexHard3(b *testing.B) { benchmarkStdIndexHard(b, "<b>hello world</b>") }
func BenchmarkIndexHard3(b *testing.B)    { benchmarkIndexHard(b, "<b>hello world</b>") }
func BenchmarkStdIndexHard4(b *testing.B) {
	benchmarkStdIndexHard(b, "<pre><b>hello</b><strong>world</strong></pre>")
}
func BenchmarkIndexHard4(b *testing.B) {
	benchmarkIndexHard(b, "<pre><b>hello</b><strong>world</strong></pre>")
}
