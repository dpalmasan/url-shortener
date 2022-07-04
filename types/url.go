package types

type Url struct {
	Id         uint64 `json:"id" bson:"id,omitempty"`
	Url        string `json:"url" bson:"url,omitempty"`
	ShortenUrl string `json:"email" bson:"shorten_url,omitempty"`
}
