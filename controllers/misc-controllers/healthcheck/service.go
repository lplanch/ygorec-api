package getHealthcheck

import (
	"time"

	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
)

type Service interface {
	GetHealthcheckService() *model.StaticHealth
}

type service struct{}

func NewServiceGet() *service { return &service{} }

func (s *service) GetHealthcheckService() *model.StaticHealth {

	healthcheck := model.StaticHealth{Status: "Available", Uptime: util.GetUptime().Milliseconds(), Date: time.Now()}

	return &healthcheck
}
