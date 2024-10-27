package usecase

import (
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener/repository"
)

type GetShortUrlUseCase struct {
	Repo repository.URLRepository
}

func (u *GetShortUrlUseCase) GetShortUrlUseCase(hash string) (string, error) {
	return u.Repo.GetLongURL(hash)
}
