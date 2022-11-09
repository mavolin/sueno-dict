package sueno

import "strings"

// IsAdjective returns true if the given word could be grammatically an
// adjective, according to its suffix.
func IsAdjective(word string) bool {
	if len(word) < 2 {
		return false
	}

	return strings.HasSuffix(word, "a") &&
		/* not a participle */ !strings.HasSuffix(word, "ata")
}

// ToAdjective returns the adjective form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToAdjective(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "a"
}
