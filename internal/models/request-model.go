package models

type UrlShortenerReq struct {
	Url    string `json:"url"`    // Url represents the original URL to be shortened.
	Length int64  `json:"length"` // Length specifies the desired length of the shortened URL.
}
