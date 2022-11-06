// Package sueno provides utilities for building sueno words.
package sueno

func RootOf(word string) string {
	switch {
	case IsAdjective(word):
		return word[:len(word)-1]
	case IsAdverb(word):
		return word[:len(word)-1]
	case IsParticiple(word):
		return word[:len(word)-3]
	case IsSingular(word):
		return word[:len(word)-1]
	case IsPlural(word):
		return word[:len(word)-2]
	case IsInfinitiveMood(word):
		return word[:len(word)-1]
	case IsImperativeMood(word):
		return word[:len(word)-2]
	case IsSubjunctiveMood(word):
		return word[:len(word)-3]
	case IsConditionalMood(word):
		return word[:len(word)-3]
	case IsPast(word), IsPresent(word), IsFuture(word): // can only be indicative
		return word[:len(word)-2]
	default:
		return ""
	}
}
