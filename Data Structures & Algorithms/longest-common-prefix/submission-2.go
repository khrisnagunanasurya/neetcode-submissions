func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]

	for pos := 0; pos < len(first); pos++ {
		for _, word := range strs[1:] {
			if pos >= len(word) || word[pos] != first[pos] {
				return first[:pos]
			}
		}
	}

	return first
}	
