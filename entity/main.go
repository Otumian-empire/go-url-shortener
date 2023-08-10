package entity

type Url struct {
	Id        int    `db:"id" json:"id"`
	Short     string `db:"short" json:"short"`
	Original  string `db:"original" json:"original"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type UrlStore interface {
	CreateUrl(shortUrl, longUrl string) (id int, err error)
	Url(id int) (urlObj Url, err error)
	Urls() (urls []Url, err error)
	DeleteUrl(id int) (err error)
}

// move this to it's own file when more entities are added
type Store interface {
	UrlStore
}
