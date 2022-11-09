package sueno

import "strings"

// IsAdverb returns true if the given word could be grammatically an adverb,
// according to its suffix.
func IsAdverb(word string) bool {
	return IsBaseAdverb(word) || IsComparativeAdverb(word) || IsSuperlativeAdverb(word)
}

// IsBaseAdverb returns true if the given word could be grammatically an adverb
// in its base form, according to its suffix.
func IsBaseAdverb(word string) bool {
	return hasGrammaticalSuffix(word, "e") &&
		/* not a participle */ !strings.HasSuffix(word, "ate")
}

// ToBaseAdverb returns the base adverb form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToBaseAdverb(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "e"
}

// IsComparativeAdverb returns true if the given word could be grammatically an
// adverb in its comparative form, according to its suffix.
func IsComparativeAdverb(word string) bool {
	return hasGrammaticalSuffix(word, "ie")
}

// ToComparativeAdverb returns the comparative adverb form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToComparativeAdverb(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ie"
}

// IsSuperlativeAdverb returns true if the given word could be grammatically an
// adverb in its superlative form, according to its suffix.
func IsSuperlativeAdverb(word string) bool {
	return hasGrammaticalSuffix(word, "oe")
}

// ToSuperlativeAdverb returns the superlative adverb form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToSuperlativeAdverb(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "oe"
}
