// Package sueno provides utilities for building sueno words.
package sueno

func RootOf(word string) string {
	switch {
	case IsBaseAdjective(word):
		return word[:len(word)-1]
	case IsComparativeAdjective(word), IsComparativeAdjective(word):
		return word[:len(word)-2]

	case IsBaseAdverb(word):
		return word[:len(word)-1]
	case IsComparativeAdverb(word), IsSuperlativeAdverb(word):
		return word[:len(word)-2]

	case IsSingular(word):
		return word[:len(word)-1]
	case IsPlural(word):

		return word[:len(word)-2]
	case IsInfinitiveMood(word):
		return word[:len(word)-1]
	case IsIndicativeMood(word):
		return word[:len(word)-2]
	case IsImperativeMood(word):
		return word[:len(word)-2]
	case IsSubjunctiveMood(word):
		return word[:len(word)-3]
	case IsConditionalMood(word):
		return word[:len(word)-3]

	case IsParticiple(word):
		return word[:len(word)-3]

	default:
		return ""
	}
}

// hasGrammaticalSuffix asserts that word...
//
//  1. ...is at least one character longer than the passed suffix.
//  2. ...has suffix as its suffix.
//  3. ...has a consonant immediately preceding the suffix.
func hasGrammaticalSuffix(word, suffix string) bool {
	if len(word) <= len(suffix) {
		return false
	}

	// a grammatical suffix must always be preceded by a consonant
	if !IsConsonant(rune(word[len(word)-len(suffix)-1])) {
		return false
	}

	return word[len(word)-len(suffix):] == suffix
}

func IsConsonant(r rune) bool {
	return !IsVowel(r)
}

func IsVowel(r rune) bool {
	switch r {
	case 'A', 'E', 'I', 'O', 'U':
		fallthrough
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
