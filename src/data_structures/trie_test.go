package data_structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_NewEmpty(t *testing.T) {
	tr := NewEmptyTrie()

	for r := 'a'; r <= 'z'; r++ {
		assert.False(t, tr.ContainsPrefix(string(r)))
	}

	for r := 'A'; r <= 'Z'; r++ {
		assert.False(t, tr.ContainsPrefix(string(r)))
	}
}

func TestTrie_NewFromArray(t *testing.T) {
	words := []string{
		"coconut",
		"raspberry",
		"blueberry",
		"peach",
	}
	tr := NewTrieFromArray(words)

	for _, w := range words {
		assert.True(t, tr.ContainsWord(w))
	}

	assert.False(t, tr.ContainsWord("not included"))
}

func TestTrie_InsertContainsWord(t *testing.T) {
	tr := NewEmptyTrie()

	tr.Insert("raspberry")
	assert.True(t, tr.ContainsWord("raspberry"))
	assert.False(t, tr.ContainsWord("raspberr"))
	assert.False(t, tr.ContainsWord("raspberrys"))

	tr.Insert("raspberr")
	assert.True(t, tr.ContainsWord("raspberr"))
	assert.False(t, tr.ContainsWord("raspber"))
	assert.True(t, tr.ContainsWord("raspberry"))

	tr.Insert("coconut in a desert")
	assert.True(t, tr.ContainsWord("coconut in a desert"))
	assert.False(t, tr.ContainsWord("coconut"))
	assert.False(t, tr.ContainsWord("coconut in a desert!"))
}

func TestTrie_InsertContainsPrefix(t *testing.T) {
	tr := NewEmptyTrie()

	assert.True(t, tr.ContainsPrefix(""))

	word := "raspberry"
	tr.Insert(word)

	for i := range word {
		assert.True(t, tr.ContainsPrefix(word[:i+1]))
	}
	assert.False(t, tr.ContainsPrefix(word+"s"))

	word2 := "raspberr"
	tr.Insert(word2)

	for i := range word {
		assert.True(t, tr.ContainsPrefix(word[:i+1]))
	}
	assert.False(t, tr.ContainsPrefix(word+"s"))

	word3 := "coconut in a desert"
	tr.Insert(word3)

	for i := range word3 {
		assert.True(t, tr.ContainsPrefix(word3[:i+1]))
	}
	assert.False(t, tr.ContainsPrefix(word3+"s"))
}
