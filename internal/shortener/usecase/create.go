package usecase

import (
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/repository"
)

type CreateShortUrlUseCase struct {
	Repo repository.URLRepository
}

func (u *CreateShortUrlUseCase) CreateShortUrl(shortUrl shortener.ShortURL) error {
	return u.Repo.CreateShortURL(shortUrl)
}
