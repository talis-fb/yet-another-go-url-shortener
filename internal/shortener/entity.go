package shortener

type Shortener struct {
	Hash    string `json:"hash"`
	LongURL string `json:"long_url"`
}
