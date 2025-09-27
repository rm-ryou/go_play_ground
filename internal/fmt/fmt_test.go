package fmt

import (
	org_fmt "fmt"
	"io"
	"os"
	"syscall"
	"testing"
)

func captureStdout(t *testing.T, f func() (int, error)) (str string, n int, e error) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("failed to create tmpfile: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	originalFd, err := syscall.Dup(STDOUT_FILENO)
	if err != nil {
		t.Fatalf("failed to dup STDOUT_FILENO: %v", err)
	}
	defer func() {
		syscall.Dup2(originalFd, STDOUT_FILENO)
		syscall.Close(originalFd)
	}()
	err = syscall.Dup2(int(tmpFile.Fd()), STDOUT_FILENO)
	if err != nil {
		t.Fatalf("failed to dup2: %v", err)
	}

	n, e = f()

	tmpFile.Seek(0, 0)
	output, err := io.ReadAll(tmpFile)
	if err != nil {
		t.Fatalf("failed to ReadAll from tmpFile: %v", err)
	}

	str = string(output)
	return
}

func TestPrintln(t *testing.T) {
	testCases := []struct {
		name  string
		input []any
	}{
		{
			name:  "文字列を標準出力に出力する",
			input: []any{"Hello, World"},
		},
		{
			name:  "output empty line. when input is empty",
			input: []any{},
		},
		{
			name:  "output '<nil>' when input is nil",
			input: []any{nil},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actStr, actN, actErr := captureStdout(t, func() (int, error) {
				return Println(tc.input...)
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
