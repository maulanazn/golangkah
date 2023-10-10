package service_test

import (
	"golangkah/contract"
	"golangkah/entity"
	"golangkah/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var bestAlbumContract = &contract.BestAlbumMock{Mock: mock.Mock{}}
var bestAlbumService = service.BestAlbumService{Contract: bestAlbumContract}

func TestBestAlbumService_GetNotFound(t *testing.T) {
	bestAlbumContract.Mock.On("FindById", "1").Return(nil)

	result, err := bestAlbumService.Get("1")

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestBestAlbumService_GetFound(t *testing.T) {
	album := entity.Album{
		Id:   "9",
		Name: "Folklore",
	}

	bestAlbumContract.Mock.On("FindById", "9").Return(album)

	result, err := bestAlbumService.Get("9")

	assert.Nil(t, err)
	assert.NotNil(t, result)
}
