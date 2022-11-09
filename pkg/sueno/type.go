package sueno

type WordType uint8

const (
	InvalidWordType WordType = iota
	Noun
	Verb
	Participle
	Adjective
	Adverb
	Preposition
	Conjunction
	Interjection
	Pronoun
	Article
	Cardinal
	Ordinal
	Abbreviation
	Other
)

// Type attempts to derive the type of the given word by matching it against
// all pronouns, articles, and numerals.
//
// If it does not match any of these, it attempts to infer the type of the word
// by analyzing its ending according to the grammar rules.
//
// Note that this solely indicates that the word could be of the returned type,
// however, it is also possible that the word is a base word that could
// theoretically be of the returned type because of its ending.
//
// If it succeeds, it returns the type.
// Otherwise, it returns InvalidWordType.
func Type(word string) WordType {
	switch {
	case IsPronoun(word):
		return Pronoun
	case IsArticle(word):
		return Article
	case IsCardinal(word):
		return Cardinal
	case IsOrdinal(word):
		return Ordinal
	case IsAdjective(word):
		return Adjective
	case IsAdverb(word):
		return Adverb
	case IsNoun(word):
		return Noun
	case IsVerb(word):
		return Verb
	case IsParticiple(word):
		return Participle
	case IsAdverb(word):
		return Adverb
	default:
		return InvalidWordType
	}
}
