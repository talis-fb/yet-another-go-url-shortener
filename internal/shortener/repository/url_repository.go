package repository

import "github.com/talis-fb/yet-another-go-url-shortener/internal/shortener"

type URLRepository interface {
	CreateShortURL(shortURL shortener.ShortURL) error
	GetLongURL(hash string) (string, error)
	DeleteShortURL(hash string) error
}
