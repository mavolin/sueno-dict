package sueno

import "strings"

func IsAdverb(word string) bool {
	if len(word) < 2 {
		return false
	}

	return strings.HasSuffix(word, "e") &&
		/* not a participle */ !strings.HasSuffix(word, "ate")
}

// ToAdverb returns the adverb form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToAdverb(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "e"
}
