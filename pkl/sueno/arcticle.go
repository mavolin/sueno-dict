package sueno

const (
	DefiniteArticle   = "e"
	IndefiniteArticle = "a"
)

// IsArticle returns true if the given word is an article.
func IsArticle(word string) bool {
	return word == DefiniteArticle || word == IndefiniteArticle
}
