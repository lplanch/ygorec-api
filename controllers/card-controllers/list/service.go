package listCards

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	ListCardsService(input *InputListCards) *[]model.ModelListCardStats
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListCardsService(input *InputListCards) *[]model.ModelListCardStats {

	resultListCards := s.repository.ListCardsRepository(input)

	return resultListCards
}
