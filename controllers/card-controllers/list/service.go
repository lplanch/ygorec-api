package listCards

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	ListCardsService(input *InputListCards) (*[]model.ModelListCardStats, string)
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListCardsService(input *InputListCards) (*[]model.ModelListCardStats, string) {

	resultListCards, errGetVersion := s.repository.ListCardsRepository(input)

	return resultListCards, errGetVersion
}
