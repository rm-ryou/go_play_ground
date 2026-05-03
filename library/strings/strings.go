package strings

func IndexByte(s string, c byte) int {
	for i := range s {
		if s[i] == c {
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

	for i := range s {
		if i+n > len(s) {
			return -1
		}

		if s[i:i+n] == substr {
			return i
		}
	}
	return -1
}

func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}
