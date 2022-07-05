package types

type Url struct {
	UrlId      int64  `json:"url_id" bson:"id,omitempty"`
	Url        string `json:"url" bson:"url,omitempty"`
	ShortenUrl string `json:"shorten_url" bson:"shorten_url,omitempty"`
}
