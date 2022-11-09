package edit

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mavolin/sueno-dict/pkg/binding"
	"github.com/mavolin/sueno-dict/pkg/sueno"
	"github.com/mavolin/sueno-dict/repository"
)

//go:generate corgi create.corgi

func (h *Handler) Create(gctx *gin.Context) {
	if err := RenderCreate(gctx.Writer); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

type CreateForm struct {
	Word string `form:"word"`
	Root string `form:"root"`

	CompoundWords []string `form:"compound_words"`

	Definitions []struct {
		Type               sueno.WordType `form:"type"`
		Translation        string         `form:"translation"`
		Definition         string         `form:"definition"`
		Example            string         `form:"example"`
		ExampleTranslation string         `form:"example_translation"`
	} `form:"definitions"`
}

func (h *Handler) ProcessCreate(gctx *gin.Context) {
	var f CreateForm

	if err := binding.Bind(gctx, &f); err != nil {
		gctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	w := repository.Word{Word: f.Word, Root: f.Root}

	if len(f.CompoundWords) > 0 {
		w.CompoundWords = make([]repository.Word, len(f.CompoundWords))
		for i, formWord := range f.CompoundWords {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			cw, err := h.repo.SearchWord(ctx, formWord)
			if err != nil {
				gctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			if cw == nil {
				gctx.Status(http.StatusBadRequest)
				return
			}

			w.CompoundWords[i] = *cw
		}
	}

	if len(f.Definitions) > 0 {
		w.Definitions = make([]repository.Definition, len(f.Definitions))
		for i, d := range f.Definitions {
			w.Definitions[i] = repository.Definition{
				Type:               d.Type,
				Translation:        d.Translation,
				Definition:         d.Definition,
				Example:            d.Example,
				ExampleTranslation: d.ExampleTranslation,
			}
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if _, err := h.repo.CreateWord(ctx, w); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gctx.Redirect(http.StatusFound, "/?q="+url.QueryEscape(w.Word))
}
