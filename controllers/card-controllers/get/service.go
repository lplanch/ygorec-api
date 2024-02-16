package getCard

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	GetCardService(input *InputServiceGetCard) (*model.EntityCard, string)
}

type service struct {
	repository Repository
}

func NewServiceGet(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCardService(input *InputServiceGetCard) (*model.EntityCard, string) {

	card := model.EntityCard{
		ID: input.ID,
	}
	resultGetCard, errGetVersion := s.repository.GetCardRepository(&card)

	return resultGetCard, errGetVersion
}
