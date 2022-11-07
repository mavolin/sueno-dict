package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mavolin/sueno-dict/handler/edit"
	"github.com/mavolin/sueno-dict/handler/root"
	"github.com/mavolin/sueno-dict/repository"
	"github.com/mavolin/sueno-dict/static"
)

type Options struct {
	Repository repository.Repository

	ProtectedUser     string
	ProtectedPassword string
}

func Setup(o Options) (*gin.Engine, error) {
	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	setupStatic(r)

	setupUnprotected(r, o)
	setupProtected(r, o)

	return r, nil
}

func setupUnprotected(r *gin.Engine, o Options) {
	g := r.Group("")

	rh := root.NewHandler(o.Repository)
	rh.RegisterRoutes(g)
}

func setupProtected(r *gin.Engine, o Options) {
	g := r.Group("")

	g.Use(gin.BasicAuth(gin.Accounts{
		o.ProtectedUser: o.ProtectedPassword,
	}))

	eh := edit.NewHandler(o.Repository)
	eh.RegisterRoutes(g)
}

func setupStatic(r *gin.Engine) {
	r.StaticFS("/static", http.FS(static.FS))
}
