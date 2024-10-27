package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/adapter"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/repository"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/usecase"
	"time"
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

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			<-ticker.C

			fmt.Println("They...")
			fmt.Println(urlRepository.FindAll())
		}
	}()

	r.GET("/shortener/:hash", shortenerHttpAdapter.Get)
	r.POST("/shortener", shortenerHttpAdapter.Create)
	r.DELETE("/shortener/:hash", shortenerHttpAdapter.Revoke)
}
