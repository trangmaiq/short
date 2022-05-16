package registry

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/trangmaiq/short/internal/handler/url"
	"github.com/trangmaiq/short/internal/persistence/cassandra"
	"github.com/trangmaiq/short/internal/services/kgs"
)

var _ Registry = new(DefaultRegistry)

type DefaultRegistry struct {
	engine         *gin.Engine
	urlRouterGroup *gin.RouterGroup

	persister *cassandra.Persister

	kgsClient url.KeyGenerator
}

func New(engine *gin.Engine) (Registry, error) {
	var (
		r   = new(DefaultRegistry)
		err = r.Init(context.TODO(), engine)
	)

	return r, err
}

func (r *DefaultRegistry) Init(_ context.Context, engine *gin.Engine) error {
	r.engine = engine
	r.urlRouterGroup = engine.Group("/urls")

	r.URLHandler().RegisterRoutes()

	persister, err := cassandra.NewPersister()
	if err != nil {
		return err
	}
	r.persister = persister

	r.kgsClient = kgs.NewClient()

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

func (r *DefaultRegistry) KeyGenerator() url.KeyGenerator {
	return r.kgsClient
}
