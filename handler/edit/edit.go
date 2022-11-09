package edit

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
	r.GET("/create", h.Create)
	r.POST("/create", h.ProcessCreate)
	r.POST("/:id/delete", h.ProcessDelete)
}
