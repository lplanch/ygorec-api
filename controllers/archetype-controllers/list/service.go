package listArchetypes

import (
	model "github.com/lplanch/test-go-api/models"
)

type Service interface {
	ListArchetypesService(input *InputListArchetypes) *[]model.ModelArchetype
}

type service struct {
	repository Repository
}

func NewServiceList(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListArchetypesService(input *InputListArchetypes) *[]model.ModelArchetype {

	resultListArchetypes := s.repository.ListArchetypesRepository(input)

	return resultListArchetypes
}
