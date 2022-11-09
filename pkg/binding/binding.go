// Package binding provides a custom binding engine base on
// github.com/go-playground/validator.
package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/form"
	"github.com/pkg/errors"
)

// decoder is the decoder singleton used for binding data.
var decoder = form.NewDecoder()

const maxMemory = 50 << 20 // 50 MB

// Bind binds the PostForm as found in the gin.Context's request to obj.
func Bind(gctx *gin.Context, dst interface{}) error {
	if err := gctx.Request.ParseMultipartForm(maxMemory); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return errors.WithStack(err)
		}
	}

	err := decoder.Decode(dst, gctx.Request.Form)
	return errors.WithStack(err)
}
