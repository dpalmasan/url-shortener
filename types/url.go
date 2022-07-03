package types

type UrlId uint32

type Url struct {
	Id         UrlId  `json:"id" bson:"id,omitempty"`
	Url        string `json:"url" bson:"url,omitempty"`
	ShortenUrl string `json:"email" bson:"shorten_url,omitempty"`
}
