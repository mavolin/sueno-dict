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

// ToOrdinal converts the given cardinal to its ordinal form.
//
// num must be a valid number.
func ToOrdinal(num string) string {
	return num + "i"
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

// ToCardinal converts the given ordinal to its cardinal form.
//
// ordinal must be a valid ordinal.
func ToCardinal(ordinal string) string {
	return ordinal[:len(ordinal)-1]
}
