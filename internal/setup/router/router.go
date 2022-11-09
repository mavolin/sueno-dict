package router

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/mavolin/sueno-dict/handler/edit"
	"github.com/mavolin/sueno-dict/handler/root"
	"github.com/mavolin/sueno-dict/internal/meta"
	"github.com/mavolin/sueno-dict/repository"
	"github.com/mavolin/sueno-dict/static"
)

type Options struct {
	Repository repository.Repository

	ProtectedUser     string
	ProtectedPassword string

	Logger *zap.SugaredLogger
}

func Setup(o Options) (*gin.Engine, error) {
	if !meta.Development {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	if err := r.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	setupStatic(r.Group(""))

	mainGroup := r.Group("")

	mainGroup.Use(ginzap.RecoveryWithZap(o.Logger.Desugar(), true))
	mainGroup.Use(ginzap.Ginzap(o.Logger.Desugar(), time.RFC3339, true))

	setupUnprotected(mainGroup, o)
	setupProtected(mainGroup, o)

	return r, nil
}

func setupUnprotected(r *gin.RouterGroup, o Options) {
	rh := root.NewHandler(o.Repository)
	rh.RegisterRoutes(r)
}

func setupProtected(r *gin.RouterGroup, o Options) {
	r.Use(gin.BasicAuth(gin.Accounts{
		o.ProtectedUser: o.ProtectedPassword,
	}))

	eh := edit.NewHandler(o.Repository)
	eh.RegisterRoutes(r)
}

func setupStatic(r *gin.RouterGroup) {
	r.StaticFS("/static", http.FS(static.FS))
}
