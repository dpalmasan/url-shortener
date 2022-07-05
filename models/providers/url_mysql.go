package providers

import (
	"time"

	"github.com/url-shortener/db/mysqldb"
	"github.com/url-shortener/types"
	"github.com/url-shortener/utils"
)

type MysqlDBUrl struct{}

func (m MysqlDBUrl) GetLongUrl(shortUrl string) (types.Url, error) {
	url := types.Url{}
	err := mysqldb.Session.Find(&url, "shorten_url = ?", shortUrl).Error
	return url, err
}

// Ids shouldn't be generated this way! but it is ok for demo purposes
func (m MysqlDBUrl) CreateShortUrl(url types.Url) (types.Url, error) {
	url.UrlId = time.Now().Unix()
	url.ShortenUrl = "https://dpurl.dev/" + utils.Base62(url.UrlId)
	result := mysqldb.Session.Create(&url)
	if result.Error != nil {
		return url, result.Error
	}
	return url, result.Error
}
