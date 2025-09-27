package fmt

import (
	"golang.org/x/sys/unix"
)

const (
	STDIN_FILENO int = iota
	STDOUT_FILENO
)

const (
	nilString = "<nil>"
)

func createBytesFromArg(arg any) []byte {
	if arg == nil {
		return []byte(nilString)
	}

	var buf []byte
	switch v := arg.(type) {
	case string:
		buf = append(buf, v...)
	}

	return buf
}

func Println(a ...any) (n int, err error) {
	var buf []byte
	for i, arg := range a {
		if i > 0 {
			buf = append(buf, ' ')
		}

		buf = append(buf, createBytesFromArg(arg)...)
	}
	buf = append(buf, '\n')

	n, err = unix.Write(STDOUT_FILENO, buf)
	return
}
