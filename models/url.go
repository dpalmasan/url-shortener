package models

import "github.com/url-shortener/types"

type UrlIface interface {
	GetLongUrl(url string) (types.Url, error)
	CreateShortUrl(url types.Url) (types.Url, error)
}
