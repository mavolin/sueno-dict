package sueno

const (
	DefiniteArticle   = "o"
	IndefiniteArticle = "a"
)

// IsArticle returns true if the given word is an article.
func IsArticle(word string) bool {
	return word == DefiniteArticle || word == IndefiniteArticle
}
