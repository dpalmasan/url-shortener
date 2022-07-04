package providers

import (
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
func (m MysqlDBUrl) CreateShortUrl(longUrl string) (types.Url, error) {
	url := types.Url{Url: longUrl, ShortenUrl: ""}
	result := mysqldb.Session.Create(&url)
	if result.Error != nil {
		return url, result.Error
	}

	// Hardcoded for demo purposes
	url.ShortenUrl = "https://dpalmaurl.dev/" + utils.Base62(url.Id)
	result = mysqldb.Session.Save(&url)
	if result.Error != nil {
		result = mysqldb.Session.Delete(&url)
	}
	return url, result.Error
}
