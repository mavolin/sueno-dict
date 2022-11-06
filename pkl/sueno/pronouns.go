package sueno

// IsPronoun returns true if the given word is a personal, possessive, or
// generalizing pronoun.
func IsPronoun(word string) bool {
	return IsPersonalPronoun(word) || IsPossessivePronoun(word) || IsGeneralizingPronoun(word)
}

// IsPersonalPronoun returns true if the given word is one of the
// personal pronouns.
func IsPersonalPronoun(word string) bool {
	switch word {
	case "mi":
		fallthrough
	case "tu":
		fallthrough
	case "se", "sa", "su", "so":
		fallthrough
	case "nu":
		fallthrough
	case "vu":
		fallthrough
	case "lu":
		return true
	default:
		return false
	}
}

// IsPossessivePronoun returns true if the given word is one of the
// possessive pronouns.
func IsPossessivePronoun(word string) bool {
	switch word {
	case "mijo":
		fallthrough
	case "tujo":
		fallthrough
	case "sejo", "sajo", "sujo", "sojo":
		fallthrough
	case "nujo":
		fallthrough
	case "vujo":
		fallthrough
	case "lujo":
		return true
	default:
		return false
	}
}

// IsGeneralizingPronoun returns true if the given word is the generalizing
// pronoun 'vi'.
func IsGeneralizingPronoun(word string) bool {
	return word == "vi"
}
