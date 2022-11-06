package root

import (
	"github.com/gin-gonic/gin"

	"github.com/mavolin/sueno-dict/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {

}

func (h *Handler) Root(gctx *gin.Context) {
	value := gctx.Query("q")
	if value != "" {
		h.searchWord(gctx)
		return
	}

	// todo: render main
}
