package searchArchetype

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	SearchArchetypeService(input *InputSearchArchetype) *[]model.ModelListArchetype
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) SearchArchetypeService(input *InputSearchArchetype) *[]model.ModelListArchetype {

	resultSearchArchetype := s.repository.SearchArchetypeRepository(input)

	return resultSearchArchetype
}
