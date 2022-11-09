package edit

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mavolin/sueno-dict/repository"
)

func (h *Handler) ProcessDelete(gctx *gin.Context) {
	id, err := repository.ParseWordID(gctx.Param("id"))
	if err != nil {
		gctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = h.repo.DeleteWord(ctx, id); err != nil {
		gctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gctx.Redirect(http.StatusFound, "/")
}
