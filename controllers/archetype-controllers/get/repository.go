package getArchetype

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetArchetypeInputServiceRepository(input *InputGetArchetype) (*InputServiceGetArchetype, string)
	GetArchetypeRepository(input *InputServiceGetArchetype) (*[]model.ModelListCardStats, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetArchetypeInputServiceRepository(input *InputGetArchetype) (*InputServiceGetArchetype, string) {

	var archetypeID uint64

	db := r.db.Model(&model.EnumArchetype{})
	errorCode := make(chan string, 1)

	getArchetype := db.Debug().Select("id").Where("LOWER(REPLACE(value, ' ', '-')) = ?", input.Value).Find(&archetypeID)

	if getArchetype.RowsAffected < 1 {
		errorCode <- "GET_ARCHETYPE_INPUT_SERVICE_NOT_FOUND_404"
		return &InputServiceGetArchetype{ID: archetypeID, Limit: input.Limit, Offset: input.Offset}, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &InputServiceGetArchetype{ID: archetypeID, Limit: input.Limit, Offset: input.Offset}, <-errorCode
}

func (r *repository) GetArchetypeRepository(input *InputServiceGetArchetype) (*[]model.ModelListCardStats, string) {

	var topArchetypeCards []model.ModelListCardStats

	db := r.db
	errorCode := make(chan string, 1)

	getArchetype := db.Debug().Raw(`
		SELECT
			b.card_id AS id,
			(
				SELECT
					name
				FROM
					entity_cards
				WHERE
					id = b.card_id
			) AS label,
			CONCAT('/cards/', CONVERT(b.card_id, char)) AS url,
			(CASE WHEN ISNULL(ban.card_id) THEN 3 ELSE ban.status END) AS limitation,
			COUNT(b.CardAmount) AS amount,
			AVG(b.CardAmount) AS average
		FROM
			(
				SELECT
					da.deck_id,
					c_d.card_id,
					COUNT(c_d.card_id) AS CardAmount
				FROM
					mv_deck_archetypes da
					LEFT JOIN graph_cards_belong_to_decks c_d ON c_d.deck_id = da.deck_id
				WHERE
					da.archetype_id = ?
					AND da.weight > 5
				GROUP BY
					da.deck_id,
					c_d.card_id
			) b
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS ban ON ban.card_id = b.card_id AND ban.banlist_id = ?
			GROUP BY b.card_id, ban.status
			ORDER BY amount DESC, average DESC, id ASC
			LIMIT ?, ?;
	`, input.ID, util.GodotEnv("LAST_BANLIST"), input.Offset, input.Limit).Find(&topArchetypeCards)

	if getArchetype.RowsAffected < 1 {
		errorCode <- "GET_ARCHETYPE_NOT_FOUND_404"
		return &topArchetypeCards, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &topArchetypeCards, <-errorCode
}
