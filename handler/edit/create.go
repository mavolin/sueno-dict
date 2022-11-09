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

type (
	CreateForm struct {
		Word string `form:"word"`
		Root string `form:"root"`

		CompoundWords []string `form:"compound_words"`

		Definitions struct {
			Type               []sueno.WordType `form:"type"`
			Translation        []string         `form:"translation"`
			Definition         []string         `form:"definition"`
			Example            []string         `form:"example"`
			ExampleTranslation []string         `form:"example_translation"`
		} `form:"definitions"`
	}
)

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

	if len(f.Definitions.Type) > 0 {
		w.Definitions = make([]repository.Definition, len(f.Definitions.Type))

		// I couldn't find any lib that allowed using slices of structs as
		// form values without having to used indexes in the input's name
		// attribute.
		// So I could either use something like jquery.repeater, or just do
		// this, which should suffice and doesn't require me to bring jquery
		// into this otherwise very simple project.
		//
		// I have a slight suspicion that no lib supports this because the
		// order of items in form value arrays is not guaranteed, but I
		// couldn't confirm this during my own testing on Chrome.
		for i := 0; i < len(w.Definitions); i++ {
			w.Definitions[i].Type = f.Definitions.Type[i]

			if i < len(f.Definitions.Translation) {
				w.Definitions[i].Translation = f.Definitions.Translation[i]
			}

			if i < len(f.Definitions.Translation) {
				w.Definitions[i].Definition = f.Definitions.Definition[i]
			}

			if i < len(f.Definitions.Example) {
				w.Definitions[i].Example = f.Definitions.Example[i]
			}

			if i < len(f.Definitions.ExampleTranslation) {
				w.Definitions[i].ExampleTranslation = f.Definitions.ExampleTranslation[i]
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
