package root

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mavolin/sueno-dict/pkg/sueno"
	"github.com/mavolin/sueno-dict/repository"
)

func (h *Handler) searchWord(gctx *gin.Context) {
	query := gctx.Query("q")
	if query == "" {
		gctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	switch gctx.Query("type") {
	case "word":
		w, err := h.repo.SearchWord(ctx, query)
		if err != nil {
			gctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if w == nil {
			h.renderNotFound(gctx)
			return
		}

		h.renderEntry(gctx, *w)
		return
	case "translation":
		ts, err := h.repo.SearchTranslation(ctx, query)
		if err != nil {
			gctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if len(ts) == 0 {
			h.renderNotFound(gctx)
			return
		}

		h.renderTranslationSearchResult(gctx, ts)
		return
	}

	// There are no ordinal and cardinal pages, just cardinal pages.
	// Since IsOrdinal and IsFraction identify numbers not just by their ending
	// but by matching the entire word, we can safely check this before doing
	// any queries.
	if sueno.IsOrdinal(strings.ToLower(query)) || sueno.IsFraction(strings.ToLower(query)) {
		query = sueno.ToCardinal(query)
	}

	// This could possibly be a base word, that looks like a word of another
	// type.
	// Search for the query as provided first, before attempting to normalize
	// it.
	w, err := h.repo.SearchWord(ctx, query)
	if err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if w != nil {
		h.renderEntry(gctx, *w)
		return
	}

	// Also, this could be a search for a translation.
	tr, err := h.repo.SearchTranslation(ctx, query)
	if err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(tr) > 0 {
		h.renderTranslationSearchResult(gctx, tr)
		return
	}

	// search yielded no results, try to normalize the query
	query = strings.ToLower(query)

	switch sueno.Type(query) {
	// maybe this was a plural? only singular nouns are in the db
	case sueno.Noun:
		// no need to search for the same query again
		if sueno.IsSingular(query) && query == gctx.Query("q") {
			h.renderNotFound(gctx)
			return
		}

		query = sueno.ToSingularNoun(query)

	// maybe this is a participle? only their infinitive verb forms are in the db
	case sueno.Participle:
		query = sueno.ToInfinitive(query)

	// maybe this is a verb? only their infinitive verb forms are in the db
	case sueno.Verb:
		// no need to search for the same query again
		if sueno.IsInfinitiveMood(query) && query == gctx.Query("q") {
			h.renderNotFound(gctx)
			return
		}

		query = sueno.ToInfinitive(query)

	// maybe this is an adjective? only their base forms are in the db
	case sueno.Adjective:
		// no need to search for the same query again
		if sueno.IsBaseAdjective(query) && query == gctx.Query("q") {
			h.renderNotFound(gctx)
			return
		}

	// maybe this is an adverb? only the adjective form of constructed adverbs is in the db
	case sueno.Adverb:
		query = sueno.ToBaseAdjective(query)
	default:
		// possibly a number 1 000 000+
		if strings.HasSuffix(query, "i") || strings.HasSuffix(query, "je") {
			query = sueno.ToCardinal(query)
		} else {
			h.renderNotFound(gctx)
			return
		}
	}

	w, err = h.repo.SearchWord(ctx, query)
	if err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if w != nil {
		h.renderEntry(gctx, *w)
		return
	}

	h.renderNotFound(gctx)
}

//go:generate corgi not_found.corgi

func (h *Handler) renderNotFound(gctx *gin.Context) {
	gctx.Status(http.StatusNotFound)

	if err := RenderNotFound(gctx.Writer, gctx.Query("q")); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

//go:generate corgi entry.corgi

func (h *Handler) renderEntry(gctx *gin.Context, w repository.Word) {
	var otherRootWords []repository.Word

	if w.Root != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		root, err := h.repo.WordRoot(ctx, w.Root)
		if err != nil {
			gctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		for i, cmp := range root.Words {
			if cmp.ID == w.ID {
				otherRootWords = append(root.Words[:i], root.Words[i+1:]...)
				break
			}
		}
	}

	if err := RenderEntry(gctx.Writer, w, otherRootWords); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	return
}

//go:generate corgi translation_search.corgi

func (h *Handler) renderTranslationSearchResult(gctx *gin.Context, ts []repository.TranslatedWord) {
	if len(ts) == 1 {
		gctx.Redirect(http.StatusFound, "/?q="+url.QueryEscape(ts[0].Word.Word)+"&type=word")
	}

	if err := RenderTranslationSearch(gctx.Writer, gctx.Query("q"), ts); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
