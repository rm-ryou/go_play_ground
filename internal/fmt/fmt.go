package fmt

import (
	"golang.org/x/sys/unix"
)

const (
	STDIN_FILENO int = iota
	STDOUT_FILENO
)

func Println(a ...any) (n int, err error) {
	var buf []byte
	for i, arg := range a {
		if i > 0 {
			buf = append(buf, ' ')
		}

		switch v := arg.(type) {
		case string:
			buf = append(buf, v...)
		}
	}
	buf = append(buf, '\n')

	n, err = unix.Write(STDOUT_FILENO, buf)
	return
}
