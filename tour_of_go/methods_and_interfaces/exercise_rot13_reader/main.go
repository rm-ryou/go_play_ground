package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	size, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i, v := range b {
		if v >= 'A' && v <= 'Z' {
			v = ((v - 'A' + 13) % 26) + 'A'
		}
		if v >= 'a' && v <= 'z' {
			v = ((v - 'a' + 13) % 26) + 'a'
		}
		b[i] = v
	}

	return size, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{
		r: s,
	}
	io.Copy(os.Stdout, &r)
}
