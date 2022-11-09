package sueno

import "strings"

// IsAdjective returns true if the given word could be grammatically an
// adjective, according to its suffix.
func IsAdjective(word string) bool {
	return IsBaseAdjective(word) || IsComparativeAdjective(word) || IsSuperlative(word)
}

// IsBaseAdjective returns true if the given word could be grammatically an
// adjective in its base form, according to its suffix.
func IsBaseAdjective(word string) bool {
	return hasGrammaticalSuffix(word, "a") && !strings.HasSuffix(word, "ata")
}

// ToBaseAdjective returns the base adjective form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToBaseAdjective(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "a"
}

// IsComparativeAdjective returns true if the given word could be grammatically
// an adjective in its comparative form, according to its suffix.
func IsComparativeAdjective(word string) bool {
	return hasGrammaticalSuffix(word, "ia")
}

// ToComparativeAdjective returns the comparative adjective form of the given
// word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToComparativeAdjective(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ia"
}

// IsSuperlative returns true if the given word could be grammatically an
// adjective in its superlative form, according to its suffix.
func IsSuperlative(word string) bool {
	return hasGrammaticalSuffix(word, "oa")
}

// ToSuperlativeAdjective returns the superlative adjective form of the given
// word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToSuperlativeAdjective(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "oa"
}
