package sueno

// IsNoun returns true if the given word could be grammatically a noun.
func IsNoun(word string) bool {
	return IsSingular(word) || IsPlural(word)
}

// IsSingular returns true if the given noun is singular.
func IsSingular(word string) bool {
	return hasGrammaticalSuffix(word, "o")
}

// ToSingularNoun returns the singular noun form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToSingularNoun(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "o"
}

// IsPlural returns true if the given noun is plural.
func IsPlural(word string) bool {
	return hasGrammaticalSuffix(word, "os")
}

// ToPluralNoun returns the plural noun form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPluralNoun(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "os"
}
