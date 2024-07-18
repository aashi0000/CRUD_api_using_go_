package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
type AlbumDB interface {
	AddAlbums(album Album) (Album, error)
	UpdateAlbum(id string,album Album) (Album, error)
	DeleteAlbum(id string) (error)
	GetAlbums() ([]Album)
	GetAlbumById(id string) (Album, error)
}