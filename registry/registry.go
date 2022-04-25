package registry

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Registry interface {
	Init(ctx context.Context, engine *gin.Engine) error
}
