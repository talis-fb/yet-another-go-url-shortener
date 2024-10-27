package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/usecase"
)

type ShortenerHttpAdapter struct {
	CreateUseCase usecase.CreateShortUrlUseCase
	RevokeUseCase usecase.RevokeShortUrlUseCase
	GetUseCase    usecase.GetShortUrlUseCase
}

func NewShortenerHttpAdapter(
	createUseCase usecase.CreateShortUrlUseCase,
	revokeUseCase usecase.RevokeShortUrlUseCase,
	getUseCase usecase.GetShortUrlUseCase,
) *ShortenerHttpAdapter {
	return &ShortenerHttpAdapter{
		CreateUseCase: createUseCase,
		RevokeUseCase: revokeUseCase,
		GetUseCase:    getUseCase,
	}
}

func (a *ShortenerHttpAdapter) Create(c *gin.Context) {
	var longUrl struct {
		LongURL string `json:"long_url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&longUrl); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrlEntity := shortener.NewShortUrlFromLongUrl(longUrl.LongURL)

	if err := a.CreateUseCase.CreateShortUrl(shortUrlEntity); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, shortUrlEntity)
}

func (a *ShortenerHttpAdapter) Revoke(c *gin.Context) {
	hash := c.Param("hash")
	if err := a.RevokeUseCase.RevokeShortUrl(hash); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(200)
}

func (a *ShortenerHttpAdapter) Get(c *gin.Context) {
	hash := c.Param("hash")
	longUrl, err := a.GetUseCase.GetShortUrlUseCase(hash)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"long_url": longUrl,
	})
}
