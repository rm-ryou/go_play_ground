package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13R rot13Reader) Read(b []byte) (int, error) {
	n, err := r13R.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i, v := range b[:n] {
		if v >= 'A' && v <= 'Z' {
			v = 'A' + ((v + 13 - 'A') % 26)
		}
		if v >= 'a' && v <= 'z' {
			v = 'a' + ((v + 13 - 'a') % 26)
		}

		b[i] = v
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
