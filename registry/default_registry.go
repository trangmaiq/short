package registry

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/trangmaiq/short/internal/handler/url"
	"github.com/trangmaiq/short/internal/persistence/inmem"
)

var _ Registry = new(DefaultRegistry)

type DefaultRegistry struct {
	engine    *gin.Engine
	persister *inmem.Persister

	urlRouterGroup *gin.RouterGroup
}

func New(engine *gin.Engine) (Registry, error) {
	var (
		r   = new(DefaultRegistry)
		err = r.Init(context.TODO(), engine)
	)

	return r, err
}

func (r *DefaultRegistry) Init(ctx context.Context, engine *gin.Engine) error {
	r.engine = engine
	r.urlRouterGroup = engine.Group("/urls")

	r.URLHandler().RegisterRoutes()

	return nil
}

func (r *DefaultRegistry) URLHandler() *url.Handler {
	return url.NewHandler(r)
}

func (r *DefaultRegistry) URLRoutes() gin.IRoutes {
	return r.urlRouterGroup
}

func (r *DefaultRegistry) URLPersister() url.Persister {
	return r.persister
}
