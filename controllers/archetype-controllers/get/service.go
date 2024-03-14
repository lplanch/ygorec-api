package getArchetype

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	GetArchetypeIDFromNameService(input *InputGetArchetype) (*InputServiceGetArchetype, string)
	GetArchetypeService(input *InputServiceGetArchetype) (*[]model.ModelListCardStats, string)
}

type service struct {
	repository Repository
}

func NewServiceGet(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetArchetypeIDFromNameService(input *InputGetArchetype) (*InputServiceGetArchetype, string) {
	resultGetArchetypeID, errGetVersion := s.repository.GetArchetypeInputServiceRepository(input)

	return resultGetArchetypeID, errGetVersion
}

func (s *service) GetArchetypeService(input *InputServiceGetArchetype) (*[]model.ModelListCardStats, string) {

	resultGetArchetype, errGetVersion := s.repository.GetArchetypeRepository(input)

	return resultGetArchetype, errGetVersion
}
