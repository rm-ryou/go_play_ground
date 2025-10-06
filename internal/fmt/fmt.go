package fmt

import (
	"syscall"
)

const (
	STDIN_FILENO int = iota
	STDOUT_FILENO
)

const (
	nilString   = "<nil>"
	trueString  = "true"
	falseString = "false"
)

func fmtInteger(u uint64) []byte {
	var buf []byte
	for u > 0 {
		rem := u / 10
		buf = append([]byte{byte('0' + u - rem*10)}, buf...)
		u = rem
	}
	return buf
}

func createBytesFromArg(arg any) []byte {
	if arg == nil {
		return []byte(nilString)
	}

	var buf []byte
	switch v := arg.(type) {
	case bool:
		if v {
			buf = append(buf, trueString...)
		} else {
			buf = append(buf, falseString...)
		}
	case string:
		buf = append(buf, v...)
	case []byte:
		buf = append(buf, '[')
		for i, b := range v {
			if i > 0 {
				buf = append(buf, ' ')
			}
			buf = append(buf, fmtInteger(uint64(b))...)
		}
		buf = append(buf, ']')
	case int32:
		buf = append(buf, fmtInteger(uint64(v))...)
	case uint8:
		buf = append(buf, fmtInteger(uint64(v))...)
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

	n, err = syscall.Write(STDOUT_FILENO, buf)
	return
}
