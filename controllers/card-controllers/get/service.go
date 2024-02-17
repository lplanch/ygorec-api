package getCard

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
)

type Service interface {
	GetCardService(input *InputServiceGetCard) (*model.ModelCard, string)
}

type service struct {
	repository Repository
}

func NewServiceGet(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCardService(input *InputServiceGetCard) (*model.ModelCard, string) {

	card := model.EntityCard{
		ID: input.ID,
	}
	resultGetCard, errGetVersion := s.repository.SanitizeCardRepository(&card)

	finalCard := model.ModelCard{
		ID:         resultGetCard.ID,
		Name:       resultGetCard.Name,
		Desc:       resultGetCard.Desc,
		Attribute:  resultGetCard.Attribute,
		Types:      util.ReallySplit(resultGetCard.Types, ","),
		Race:       resultGetCard.Race,
		Archetypes: util.ReallySplit(resultGetCard.Archetypes, ","),
		Atk:        resultGetCard.Atk,
		Def:        resultGetCard.Def,
		Level:      resultGetCard.Level,
		Categories: util.ReallySplit(resultGetCard.Categories, ","),
	}

	return &finalCard, errGetVersion
}
