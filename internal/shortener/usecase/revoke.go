package usecase

import (
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/repository"
)

type RevokeShortUrlUseCase struct {
	Repo repository.URLRepository
}

func (u *RevokeShortUrlUseCase) RevokeShortUrl(hash string) error {
	return u.Repo.DeleteShortURL(hash)
}
