package strings

func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}

func ContainsAny(s, chars string) bool {
	return IndexAny(s, chars) >= 0
}
