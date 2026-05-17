package strings

import "unsafe"

func Clone(s string) string {
	if s == "" {
		return ""
	}

	b := make([]byte, len(s))
	copy(b, s)
	// 不要なコピーをなくし、bしか返さないため、unsafeでもない
	return unsafe.String(&b[0], len(b))
}
