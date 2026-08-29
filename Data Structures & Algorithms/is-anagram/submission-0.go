func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a, b := dict(s), dict(t)

	for char, count := range a {
		if b[char] != count {
			return false
		}
	}

	return true
}

func dict(s string) map[rune]int {
	chars := make(map[rune]int)

	for _, char := range s {
		chars[char]++
	}	

	return chars
}