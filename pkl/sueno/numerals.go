package sueno

// IsNumeral reports whether the given word represents a number.
//
// Multiple words will always yield false, even if they together still form a
// number.
func IsNumeral(word string) bool {
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
	case "fin", "findez", "finsent", "finmil":
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
