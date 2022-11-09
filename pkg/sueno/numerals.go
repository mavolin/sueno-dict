package sueno

// IsCardinal reports whether the given word represents a cardinal.
//
// Multiple words will always yield false, even if they together still form a
// number.
func IsCardinal(word string) bool {
	switch word {
	case "nul":
		fallthrough
	case "uno", "unodez", "unosent", "unomil":
		fallthrough
	case "duo", "duodez", "duosent", "duomil":
		fallthrough
	case "tre", "tredez", "tresent", "tremil":
		fallthrough
	case "var", "vardez", "varsent", "varmil":
		fallthrough
	case "fun", "fundez", "funsent", "funmil":
		fallthrough
	case "son", "sondez", "sonsent", "sonmil":
		fallthrough
	case "sep", "sepdez", "sepsent", "sepmil":
		fallthrough
	case "oka", "okadez", "okasent", "okamil":
		fallthrough
	case "nin", "nindez", "ninsent", "ninmil":
		return true
	default:
		return false
	}
}

// ToCardinal converts the given ordinal to its cardinal form.
//
// word must be a valid ordinal or denominator.
func ToCardinal(word string) string {
	return word[:len(word)-1]
}

func IsOrdinal(word string) bool {
	switch word {
	case "nuli":
		fallthrough
	case "unoi", "unodezi", "unosenti", "unomili":
		fallthrough
	case "duoi", "duodezi", "duosenti", "duomili":
		fallthrough
	case "trei", "tredezi", "tresenti", "tremili":
		fallthrough
	case "vari", "vardezi", "varsenti", "varmili":
		fallthrough
	case "funi", "fundezi", "funsenti", "funmili":
		fallthrough
	case "soni", "sondezi", "sonsenti", "sonmili":
		fallthrough
	case "sepi", "sepdezi", "sepsenti", "sepmili":
		fallthrough
	case "okai", "okadezi", "okasenti", "okamili":
		fallthrough
	case "nini", "nindezi", "ninsenti", "ninmili":
		return true
	default:
		return false
	}
}

// ToOrdinal converts the given cardinal to its ordinal form.
//
// word must be a valid cardinal or fraction.
func ToOrdinal(word string) string {
	if IsCardinal(word) {
		return word + "i"
	}

	return word[:len(word)-1] + "i"
}

// IsFraction reports whether the given word represents a fraction.
func IsFraction(word string) bool {
	switch word {
	case "nulu":
		fallthrough
	case "unou", "unodezu", "unosentu", "unomilu":
		fallthrough
	case "duou", "duodezu", "duosentu", "duomilu":
		fallthrough
	case "treu", "tredezu", "tresentu", "tremilu":
		fallthrough
	case "varu", "vardezu", "varsentu", "varmilu":
		fallthrough
	case "funu", "fundezu", "funsentu", "funmilu":
		fallthrough
	case "sonu", "sondezu", "sonsentu", "sonmilu":
		fallthrough
	case "sepu", "sepdezu", "sepsentu", "sepmilu":
		fallthrough
	case "okau", "okadezu", "okasentu", "okamilu":
		fallthrough
	case "niniu", "nindezu", "ninsentu", "ninmilu":
		return true
	default:
		return false
	}
}

// ToFraction converts the given cardinal to its fraction form.
//
// word must be a valid cardinal or ordinal.
func ToFraction(word string) string {
	if IsCardinal(word) {
		return word + "u"
	}

	return word[:len(word)-1] + "u"
}
