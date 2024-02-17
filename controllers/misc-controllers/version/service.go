package getVersion

import (
	"slices"
	"time"

	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
)

type Service interface {
	GetVersionService() (*model.StaticVersion, string)
}

type service struct {
	repository Repository
}

func NewServiceGet(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetVersionService() (*model.StaticVersion, string) {

	var getVersion model.StaticVersion
	listKeyValue, errGetVersion := s.repository.GetVersionRepository()

	i := slices.IndexFunc[[]model.KeyValueStore](*listKeyValue, func(c model.KeyValueStore) bool { return c.Key == util.KV_ENUM_LAST_COMMIT })
	j := slices.IndexFunc[[]model.KeyValueStore](*listKeyValue, func(c model.KeyValueStore) bool { return c.Key == util.KV_BABELCDB_LAST_COMMIT })
	k := slices.IndexFunc[[]model.KeyValueStore](*listKeyValue, func(c model.KeyValueStore) bool { return c.Key == util.KV_VERSION_DATE })

	if i < 0 || j < 0 || k < 0 {
		errGetVersion = "VERSIONS_NOT_FOUND_500"
		return &getVersion, errGetVersion
	}
	getVersion.EnumLastCommit = (*listKeyValue)[i].Value
	getVersion.CardsLastCommit = (*listKeyValue)[j].Value
	getVersion.LastUpdate, _ = time.Parse(time.RFC3339, (*listKeyValue)[k].Value)
	getVersion.CardsAmount = 10

	return &getVersion, errGetVersion
}
