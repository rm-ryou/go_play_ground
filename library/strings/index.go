package strings

import "unicode/utf8"

func IndexByte(s string, c byte) int {
	for i := range len(s) {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func IndexRune(s string, r rune) int {
	switch {
	// utf8.RuneSelf: 0x80 i.e. 1バイトで表現可能なコードポイントの時
	case r >= 0 && r < utf8.RuneSelf:
		return IndexByte(s, byte(r))
	case r == utf8.RuneError:
		for i, r := range s {
			if r == utf8.RuneError {
				return i
			}
		}
		return -1
	case !utf8.ValidRune(r):
		return -1
	default:
		for i, rr := range s {
			if rr == r {
				return i
			}
		}
	}

	return -1
}

func IndexAny(s, chars string) int {
	if chars == "" {
		return -1
	}

	for i, c := range s {
		if IndexRune(chars, c) >= 0 {
			return i
		}
	}

	return -1
}

func Index(s, substr string) int {
	n := len(substr)
	switch {
	case n == 0 || substr == s:
		return 0
	case n == 1:
		return IndexByte(s, substr[0])
	case n >= len(s):
		return -1
	}

	c0 := substr[0]
	t := len(s) - n + 1
	for i := 0; i < t; {
		if s[i] != c0 {
			o := IndexByte(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}

		if s[i:i+n] == substr {
			return i
		}

		i++
	}
	return -1
}
