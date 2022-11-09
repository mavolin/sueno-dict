package sueno

import "strings"

// IsVerb returns true if the given word is a verb in any conjugated form.
func IsVerb(word string) bool {
	return IsInfinitiveMood(word) ||
		IsPast(word) || IsPresent(word) || IsFuture(word) ||
		IsImperativeMood(word)
}

// IsPast returns true if the given word is in the past tense in any mood.
//
// The word may be in any grammatical mood.
func IsPast(word string) bool {
	return hasGrammaticalSuffix(word, "as") || hasGrammaticalSuffix(word, "uas") || hasGrammaticalSuffix(word, "oas")
}

// IsPresent returns true if the given word is in the present tense.
//
// The word may be in any grammatical mood.
func IsPresent(word string) bool {
	return hasGrammaticalSuffix(word, "es") || hasGrammaticalSuffix(word, "ues") || hasGrammaticalSuffix(word, "oes")
}

// IsFuture returns true if the given word is in the future tense.
//
// The word may be in any grammatical mood.
func IsFuture(word string) bool {
	return hasGrammaticalSuffix(word, "is") || hasGrammaticalSuffix(word, "uis") || hasGrammaticalSuffix(word, "ois")
}

// ============================================================================
// Moods
// ======================================================================================

// ===================================== Infinitive =====================================

// IsInfinitiveMood returns true if the given word is a verb in its infinitive
// mood.
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

// ===================================== Indicative =====================================

// IsIndicativeMood returns true if the given word is in the indicative mood in
// any tense.
func IsIndicativeMood(word string) bool {
	return hasGrammaticalSuffix(word, "as") ||
		hasGrammaticalSuffix(word, "es") || hasGrammaticalSuffix(word, "is")
}

// ToPastIndicative returns the past indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPastIndicative(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "as"
}

// ToPresentIndicative returns the present indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToPresentIndicative(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "es"
}

// ToFutureIndicative returns the future indicative form of the given word.
//
// The word may be a noun, a verb, a participle, an adjective, or a constructed
// adverb.
//
// If the root of the word cannot be extracted, an empty string is returned.
//
// Note that the returned word is grammatically valid, but may not exist as a
// word in the language.
func ToFutureIndicative(word string) string {
	root := RootOf(word)
	if root == "" {
		return ""
	}

	return root + "is"
}

// ===================================== Imperative =====================================

func IsImperativeMood(word string) bool {
	return hasGrammaticalSuffix(word, "us")
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

// ==================================== Subjunctive =====================================

func IsSubjunctiveMood(word string) bool {
	return hasGrammaticalSuffix(word, "uas") || hasGrammaticalSuffix(word, "ues") || hasGrammaticalSuffix(word, "uis")
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

// ==================================== Conditional =====================================

// IsConditionalMood returns true if the given word is in the conditional mood.
func IsConditionalMood(word string) bool {
	return hasGrammaticalSuffix(word, "oas") || hasGrammaticalSuffix(word, "oes") || hasGrammaticalSuffix(word, "ois")
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
