package sueno

import "strings"

// IsVerb returns true if the given word is a verb in any conjugated form.
func IsVerb(word string) bool {
	return IsInfinitiveMood(word) ||
		IsPast(word) || IsPresent(word) || IsFuture(word) ||
		IsImperativeMood(word)
}

// IsPast returns true if the given word is in the past tense.
//
// The word may be in any grammatical mood.
func IsPast(word string) bool {
	if len(word) < 3 {
		return false
	}

	return strings.HasSuffix(word, "as")
}

// ToPast returns the past indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPast(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "as"
}

// IsPresent returns true if the given word is in the present tense.
//
// The word may be in any grammatical mood.
func IsPresent(word string) bool {
	if len(word) < 3 {
		return false
	}

	return strings.HasSuffix(word, "es")
}

// ToPresent returns the present indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPresent(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "es"
}

// IsFuture returns true if the given word is in the future tense.
//
// The word may be in any grammatical mood.
func IsFuture(word string) bool {
	if len(word) < 3 {
		return false
	}

	return strings.HasSuffix(word, "is")
}

// ToFuture returns the future indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToFuture(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "is"
}

func IsInfinitiveMood(word string) bool {
	if len(word) < 2 {
		return false
	}

	return strings.HasSuffix(word, "i")
}

// ToInfinitive returns the infinitive form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToInfinitive(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "i"
}

func IsIndicativeMood(word string) bool {
	if len(word) < 3 {
		return false
	}

	if !IsPresent(word) && !IsPresent(word) && !IsFuture(word) {
		return false
	}

	if len(word) == 3 {
		return true
	}

	suffix3rd := word[len(word)-3]
	return suffix3rd != 'u' && suffix3rd != 'i'
}

func IsImperativeMood(word string) bool {
	if len(word) < 3 {
		return false
	}

	return strings.HasSuffix(word, "us")
}

// ToImperative returns the imperative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToImperative(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "us"
}

func IsSubjunctiveMood(word string) bool {
	if len(word) < 4 {
		return false
	}

	suffix3rd := word[len(word)-3]

	return suffix3rd == 'u' &&
		(IsPast(word) || IsPresent(word) || IsFuture(word))
}

// ToPastSubjunctive returns the past subjunctive form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPastSubjunctive(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "uas"
}

// ToPresentSubjunctive returns the present subjunctive form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPresentSubjunctive(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ues"
}

// ToFutureSubjunctive returns the future subjunctive form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToFutureSubjunctive(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "uis"
}

// IsConditionalMood returns true if the given word is in the conditional mood.
func IsConditionalMood(word string) bool {
	if len(word) < 4 {
		return false
	}

	suffix3rd := word[len(word)-3]

	return suffix3rd == 'o' && IsImperativeMood(word)
}

// ToPastConditional returns the past conditional form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPastConditional(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "oas"
}

// ToPresentConditional returns the present conditional form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPresentConditional(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "oes"
}

// ToFutureConditional returns the future conditional form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToFutureConditional(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "ois"
}
