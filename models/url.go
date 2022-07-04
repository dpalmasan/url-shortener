package models

import "github.com/url-shortener/types"

type UrlModel interface {
	GetLongUrl(url string) (types.Url, error)
	CreateShortUrl(url string) (types.Url, error)
}
