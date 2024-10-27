package redirect

import (
	"github.com/gin-gonic/gin"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/usecase"
	"net/http"
)

type RedirectHttpAdapter struct {
	GetUseCase usecase.GetShortUrlUseCase
}

func (a *RedirectHttpAdapter) Redirect(c *gin.Context) {
	hash := c.Param("uri")

	longUrl, err := a.GetUseCase.GetShortUrlUseCase(hash)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longUrl)
}
