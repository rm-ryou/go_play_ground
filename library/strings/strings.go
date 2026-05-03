package strings

func IndexByte(s string, c byte) int {
	for i := range len(s) {
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

func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}
