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
	if IsFraction(word) {
		return word[:len(word)-1] + "i"
	}

	return word + "i"
}

// IsFraction reports whether the given word represents a fraction.
func IsFraction(word string) bool {
	switch word {
	case "nulje":
		fallthrough
	case "unoje", "unodezje", "unosentje", "unomilje":
		fallthrough
	case "duoje", "duodezje", "duosentje", "duomilje":
		fallthrough
	case "treje", "tredezje", "tresentje", "tremilje":
		fallthrough
	case "varje", "vardezje", "varsentje", "varmilje":
		fallthrough
	case "funje", "fundezje", "funsentje", "funmilje":
		fallthrough
	case "sonje", "sondezje", "sonsentje", "sonmilje":
		fallthrough
	case "sepje", "sepdezje", "sepsentje", "sepmilje":
		fallthrough
	case "okaje", "okadezje", "okasentje", "okamilje":
		fallthrough
	case "ninje", "nindezje", "ninsentje", "ninmilje":
		return true
	default:
		return false
	}
}

// ToFraction converts the given cardinal to its fraction form.
//
// word must be a valid cardinal or ordinal.
func ToFraction(word string) string {
	if IsOrdinal(word) {
		return word[:len(word)-1] + "je"
	}

	return word + "je"
}
