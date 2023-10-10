package contract

import (
	"golangkah/entity"

	"github.com/stretchr/testify/mock"
)

type BestAlbumMock struct {
	Mock mock.Mock
}

func (bestAlbumMock *BestAlbumMock) FindById(id string) *entity.Album {
	resultMock := bestAlbumMock.Mock.Called(id)

	if resultMock.Get(0) == nil {
		return nil
	} else {
		convResult := resultMock.Get(0).(entity.Album)

		return &convResult
	}
}
