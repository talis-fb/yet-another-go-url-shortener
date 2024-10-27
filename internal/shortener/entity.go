package shortener

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type ShortURL struct {
	CreatedAt int64  `json:"created_at"`
	Hash      string `json:"hash" binding:"required"`
	LongURL   string `json:"long_url" binding:"required"`
}

func NewShortUrlFromLongUrl(longUrl string) ShortURL {
	createdAt := time.Now().Unix()
	combined := strconv.FormatInt(createdAt, 10) + longUrl
	hashBytes := sha256.Sum256([]byte(combined))
	hash := hex.EncodeToString(hashBytes[:3])

	return ShortURL{
		CreatedAt: createdAt,
		Hash:      hash,
		LongURL:   longUrl,
	}
}
