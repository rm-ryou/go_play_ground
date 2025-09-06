package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r rot13Reader) Read(b []byte) (int, error) {
	n, err := r13r.r.Read(b)
	if err != nil {
		return 0, err
	}

	for i, v := range b[:n] {
		if v >= 'A' && v <= 'Z' {
			// v += 13
			v = ((v-'A'+13)%26 + 'A')
		} else if v >= 'a' && v <= 'z' {
			// v += 13
			v = ((v-'a'+13)%26 + 'a')
		}
		b[i] = v
	}
	// fmt.Println(b[:n])
	// fmt.Printf("%q", b[:n])
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
