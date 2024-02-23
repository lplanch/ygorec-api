package listBanlists

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	ListBanlistsService() *[]model.ModelBanlist
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListBanlistsService() *[]model.ModelBanlist {

	resultListBanlists := s.repository.ListBanlistsRepository()

	return resultListBanlists
}
