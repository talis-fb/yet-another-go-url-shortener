package internal

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	redirectAdapter "github.com/talis-fb/yet-another-go-url-shortener/internal/redirect/adapter"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/adapter"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/repository"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/usecase"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	urlRepository := repository.NewURLRepositoryInMemory()
	createUseCase := usecase.CreateShortUrlUseCase{
		Repo: urlRepository,
	}

	revokeUseCase := usecase.RevokeShortUrlUseCase{
		Repo: urlRepository,
	}

	getUseCase := usecase.GetShortUrlUseCase{
		Repo: urlRepository,
	}

	shortenerHttpAdapter := adapter.NewShortenerHttpAdapter(createUseCase, revokeUseCase, getUseCase)
	redirectHttpAdapter := redirectAdapter.RedirectHttpAdapter{GetUseCase: getUseCase}

	r.GET("/shortener/:hash", shortenerHttpAdapter.Get)
	r.POST("/shortener", shortenerHttpAdapter.Create)
	r.DELETE("/shortener/:hash", shortenerHttpAdapter.Revoke)

	r.GET("/r/:uri", redirectHttpAdapter.Redirect)

	// -------------------------------------

	// For DEBUG
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C

			fmt.Println("They...")
			fmt.Println(urlRepository.FindAll())
		}
	}()
}
