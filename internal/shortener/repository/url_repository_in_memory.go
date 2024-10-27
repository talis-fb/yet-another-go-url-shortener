package repository

import (
	"errors"
	"github.com/talis-fb/yet-another-go-url-shortener/internal/shortener"
)

type URLRepositoryInMemory struct {
	urls map[string]shortener.ShortURL
}

func NewURLRepositoryInMemory() *URLRepositoryInMemory {
	return &URLRepositoryInMemory{
		urls: make(map[string]shortener.ShortURL),
	}
}

func (r *URLRepositoryInMemory) CreateShortURL(shortURL shortener.ShortURL) error {
	r.urls[shortURL.Hash] = shortURL
	return nil
}

func (r *URLRepositoryInMemory) GetLongURL(hash string) (string, error) {
	if shortURL, ok := r.urls[hash]; ok {
		return shortURL.LongURL, nil
	}
	return "", errors.New("short url not found")
}

func (r *URLRepositoryInMemory) FindAll() []shortener.ShortURL {
	values := make([]shortener.ShortURL, 0)
	for _, value := range r.urls {
		values = append(values, value)
	}
	return values
}

func (r *URLRepositoryInMemory) DeleteShortURL(hash string) error {
	delete(r.urls, hash)
	return nil
}
