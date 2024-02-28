package listCards

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	ListCardsService(input *InputListCards) *model.ModelFullListCardStats
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListCardsService(input *InputListCards) *model.ModelFullListCardStats {

	resultFullListCards := &model.ModelFullListCardStats{
		DeckAmount: *s.repository.GetDeckAmount(input),
		List:       *s.repository.ListCardsRepository(input),
	}

	return resultFullListCards
}
