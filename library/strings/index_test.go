package strings

import (
	"math/rand"
	stdStrings "strings"
	"testing"
)

const BenchmarkString = "some_text=some☺value"

var benchInputHard = makeBenchInputHard()

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
