package sueno

func IsParticiple(word string) bool {
	return IsPastParticiple(word) || IsPresentParticiple(word) || IsFutureParticiple(word)
}

func IsPastParticiple(word string) bool {
	return hasGrammaticalSuffix(word, "ata")
}

// ToPastParticiple returns the past participle form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPastParticiple(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ata"
}

func IsPresentParticiple(word string) bool {
	return hasGrammaticalSuffix(word, "ate")
}

// ToPresentParticiple returns the present participle form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPresentParticiple(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ate"
}

func IsFutureParticiple(word string) bool {
	return hasGrammaticalSuffix(word, "ati")
}

// ToFutureParticiple returns the future participle form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToFutureParticiple(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ati"
}
