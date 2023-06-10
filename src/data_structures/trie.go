package data_structures

/////////////////// Trie Definition ///////////////////

// Trie is a data structure that stores multiple strings and allows to find substrings and string for O(n) time, O(1) space
type Trie struct {
	children map[rune]*Trie
	isEnd    bool
}

// NewEmptyTrie works in O(1) time, O(1) space
func NewEmptyTrie() *Trie {
	return NewTrieFromArray(make([]string, 0))
}

// NewTrieFromArray works in O(n*m) time, O(n*m) space, where n is the words count, and m is the length of the longest word
func NewTrieFromArray(words []string) *Trie {
	t := Trie{
		children: make(map[rune]*Trie),
		isEnd:    false,
	}

	for _, word := range words {
		t.Insert(word)
	}

	return &t
}

// Insert works in O(m) time, O(m) space
func (t *Trie) Insert(word string) {
	cur := t

	for _, char := range word {
		if _, ok := cur.children[char]; !ok {
			cur.children[char] = NewEmptyTrie()
		}
		cur = cur.children[char]
	}

	cur.isEnd = true
}

// ContainsWord works in O(m) time, O(m) space
func (t *Trie) ContainsWord(word string) bool {
	cur := t

	for _, char := range word {
		if _, ok := cur.children[char]; !ok {
			return false
		}
		cur = cur.children[char]
	}

	return cur.isEnd
}

// ContainsPrefix works in O(m) time, O(m) space
func (t *Trie) ContainsPrefix(prefix string) bool {
	cur := t

	for _, char := range prefix {
		if _, ok := cur.children[char]; !ok {
			return false
		}
		cur = cur.children[char]
	}

	return true
}
