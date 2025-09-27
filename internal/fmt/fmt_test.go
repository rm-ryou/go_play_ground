package fmt

import (
	"bytes"
	org_fmt "fmt"
	"io"
	"os"
	"testing"

	"golang.org/x/sys/unix"
)

func captureStdout(t *testing.T, f func() (int, error)) (string, int, error) {
	t.Helper()

	clonedStdoutFd, err := unix.Dup(STDOUT_FILENO)
	if err != nil {
		t.Fatalf("failed to dup STDOUT_FILENO: %v", err)
	}
	defer unix.Close(clonedStdoutFd)

	backupStdout := os.Stdout
	defer func() { os.Stdout = backupStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to pipe: %v", err)
	}
	writeFd := int(w.Fd())

	if err := unix.Dup2(writeFd, clonedStdoutFd); err != nil {
		t.Fatalf("failed to dup2: %v", err)
	}

	os.Stdout = w

	outCh := make(chan string)
	go func() {
		var buf bytes.Buffer

		io.Copy(&buf, r)
		outCh <- buf.String()
	}()

	n, err := f()

	w.Close()
	r.Close()

	str := <-outCh

	return str, n, err
}

func TestPrintln(t *testing.T) {
	stringTestCases := []struct {
		name  string
		input []any
	}{
		{
			name:  "文字列を標準出力に出力する",
			input: []any{"Hello, World"},
		},
		{
			name:  "引数がない時は改行のみを出力する",
			input: []any{},
		},
	}

	for _, tc := range stringTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actStr, actN, actErr := captureStdout(t, func() (int, error) {
				// return Println(tc.input...)
				return org_fmt.Println(tc.input...)
			})

			str, n, err := captureStdout(t, func() (int, error) {
				return org_fmt.Println(tc.input...)
			})

			if actStr != str {
				t.Errorf("expected string is %v. But actual is %v", str, actStr)
			}

			if actN != n {
				t.Errorf("expected len is %v. But actual is %v", n, actN)
			}

			if actErr != err {
				t.Errorf("expected error is %v. But actual is %v", err, actErr)
			}
		})
	}
}
