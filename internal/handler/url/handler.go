package url

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	URL struct {
		ID        uuid.UUID `json:"id,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`

		BaseURL
	}

	BaseURL struct {
		OriginalURL string    `json:"original_url,omitempty"`
		Alias       string    `json:"alias,omitempty"`
		ExpiredAt   time.Time `json:"expired_at,omitempty"`

		// UTM Parameters (optional)
		// https://en.wikipedia.org/wiki/UTM_parameters
		Source   string `json:"source,omitempty"`
		Medium   string `json:"medium,omitempty"`
		Campaign string `json:"campaign,omitempty"`
	}
	CreateURLRequest BaseURL
	Persister        interface {
		CreateURL(u *URL) error
	}
	PersistenceProvider interface {
		URLPersister() Persister
	}
	handlerDependencies interface {
		URLRoutes() gin.IRoutes
		PersistenceProvider
	}
	HandlerProvider interface {
		URLHandler() *Handler
	}
	Handler struct {
		hd handlerDependencies
	}
)

func NewHandler(hd handlerDependencies) *Handler {
	return &Handler{hd}
}

func (h *Handler) RegisterRoutes() {
	h.hd.URLRoutes().POST("/", h.createURL)
}

func (r *CreateURLRequest) validate() error {
	if r == nil {
		return errors.New("request should not be empty")
	}

	if r.OriginalURL == "" {
		return errors.New("original url should not be empty")
	}

	return nil
}

func (h *Handler) createURL(c *gin.Context) {
	var request CreateURLRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Printf(fmt.Errorf("bind json create url request failed: %w", err).Error())
	}

	err = request.validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "invalid_request",
			"message":    err.Error(),
		})

		return
	}

	_, err = url.Parse(request.OriginalURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "invalid_request",
			"message":    "parse original url failed",
		})

		return
	}

	now := time.Now()
	err = h.hd.URLPersister().CreateURL(&URL{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,

		BaseURL: BaseURL(request),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_code": "internal_error",
			"message":    "create original url failed",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_url": "TVDah8nCj",
	})
}
