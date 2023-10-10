package contract

import "golangkah/entity"

type AlbumInterface interface {
	FindById(id string) *entity.Album
}
