package strings

import "unsafe"

type Builder struct {
	addr *Builder
	buf  []byte
}

// コピーされ、内部のバッファが共有されるとunsafeになるためこれを防いでいる
// これは、String()の際に、unsafeを使用せず、コピーすれば効率は落ちるが、解決する
func (b *Builder) copyCheck() {
	if b.addr == nil {
		b.addr = b
		return
	} else if b.addr != b {
		panic("strings: illegal use of non-zero Builder copied by value")
	}
}

func (b *Builder) Cap() int {
	return cap(b.buf)
}

func (b *Builder) Grow(n int) {}

func (b *Builder) Len() int {
	return len(b.buf)
}

func (b *Builder) Reset() {
	b.addr = nil
	b.buf = nil
}

func (b *Builder) String() string {
	return unsafe.String(unsafe.SliceData(b.buf), len(b.buf))
}

// func (b *Builder) Write(p []byte) (int, error)
//
// func (b *Builder) WriteByte(c byte) error
//
// func (b *Builder) WriteRune(r rune) (int, error)

func (b *Builder) WriteString(s string) (int, error) {
	b.copyCheck()
	b.buf = append(b.buf, s...)
	return len(s), nil
}
