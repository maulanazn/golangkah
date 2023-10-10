package service

import (
	"errors"
	"golangkah/contract"
	"golangkah/entity"
)

type BestAlbumService struct {
	Contract contract.AlbumInterface
}

func (bestAlbumService BestAlbumService) Get(id string) (*entity.Album, error) {
	result := bestAlbumService.Contract.FindById(id)
	if result == nil {
		return result, errors.New("Album not found")
	} else {
		return result, nil
	}
}
