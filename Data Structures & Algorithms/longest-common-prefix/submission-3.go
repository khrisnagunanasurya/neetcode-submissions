type TrieNode struct {
	children 	[26]*TrieNode
	isWord 		bool
}

func longestCommonPrefix(strs []string) string {
	root := &TrieNode{}

	for _, word := range(strs) {
		node := root

		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'

			if node.children[index] == nil {
				node.children[index] = &TrieNode{}
			}

			node = node.children[index]
		}

		node.isWord = true
	}

	node := root
	prefix := make([]byte, 0)

	for !node.isWord {
		childCount := 0

		var next *TrieNode
		var character byte

		for i, child := range node.children {
			if child != nil {
				childCount++
				next = child
				character = byte(i) + 'a'
			}
		}

		if childCount != 1 {
			break
		}

		prefix = append(prefix, character)
		node = next
	}

	return string(prefix)
}	
