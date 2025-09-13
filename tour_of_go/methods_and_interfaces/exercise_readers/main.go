package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	size := len(b)
	for i := range size {
		b[i] = 'A'
	}
	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}
