package getVersion

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	GetVersionService() (*model.StaticVersion, string)
}

type service struct {
	repository Repository
}

func NewServiceGet(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetVersionService() (*model.StaticVersion, string) {

	resultGetVersion, errGetVersion := s.repository.GetVersionRepository()

	return resultGetVersion, errGetVersion
}
