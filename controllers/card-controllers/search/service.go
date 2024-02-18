package searchCard

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	SearchCardService(input *InputSearchCard) *[]model.ModelListCard
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) SearchCardService(input *InputSearchCard) *[]model.ModelListCard {

	resultSearchCard := s.repository.SearchCardRepository(input)

	return resultSearchCard
}
