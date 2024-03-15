package getArchetype

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetArchetypeDeckAmount(input *InputServiceGetArchetype) *uint32
	GetArchetypeInputServiceRepository(input *InputGetArchetype) (*InputServiceGetArchetype, string)
	GetArchetypeCardsRepository(input *InputServiceGetArchetype) *[]model.ModelListCardStats
	GetArchetypeOtherCardsRepository(input *InputServiceGetArchetype) *[]model.ModelListCardStats
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetArchetypeDeckAmount(input *InputServiceGetArchetype) *uint32 {

	var total uint32

	db := r.db.Model(&model.MvDeckArchetypes{})

	db.Debug().Select(`
		COUNT(*)
	`).Where(`
		archetype_id = ? AND weight > 5
	`, input.ID).Find(&total)

	return &total
}

func (r *repository) GetArchetypeInputServiceRepository(input *InputGetArchetype) (*InputServiceGetArchetype, string) {

	var archetypeID uint64

	db := r.db.Model(&model.EnumArchetype{})
	errorCode := make(chan string, 1)

	getArchetype := db.Debug().Select("id").Where("LOWER(REPLACE(value, ' ', '-')) = ?", input.Value).Find(&archetypeID)

	if getArchetype.RowsAffected < 1 {
		errorCode <- "GET_ARCHETYPE_INPUT_SERVICE_NOT_FOUND_404"
		return &InputServiceGetArchetype{ID: archetypeID}, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &InputServiceGetArchetype{ID: archetypeID}, <-errorCode
}

func (r *repository) GetArchetypeCardsRepository(input *InputServiceGetArchetype) *[]model.ModelListCardStats {

	var topArchetypeCards []model.ModelListCardStats

	db := r.db

	db.Debug().Raw(`
		SELECT
			mv.card_id AS id,
			ec.name AS label,
			CONCAT('/cards/', CONVERT(mv.card_id, char)) AS url,
			(CASE WHEN ISNULL(ban.card_id) THEN 3 ELSE ban.status END) AS limitation,
			mv.deck_amount AS amount,
			mv.card_amount / mv.deck_amount AS average
		FROM mv_top_archetype_cards mv
		JOIN entity_cards ec ON ec.id = mv.card_id
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS ban ON ban.card_id = mv.card_id AND ban.banlist_id = ?
			WHERE mv.archetype_id = ? AND match_archetype(ec.set_code, mv.archetype_id)
			ORDER BY mv.deck_amount DESC, mv.card_amount DESC, mv.card_id ASC;
	`, util.GodotEnv("LAST_BANLIST"), input.ID).Find(&topArchetypeCards)

	return &topArchetypeCards
}

func (r *repository) GetArchetypeOtherCardsRepository(input *InputServiceGetArchetype) *[]model.ModelListCardStats {

	var topArchetypeCards []model.ModelListCardStats

	db := r.db

	db.Debug().Raw(`
		SELECT
			mv.card_id AS id,
			(
				SELECT
					name
				FROM
					entity_cards
				WHERE
					id = mv.card_id
			) AS label,
			CONCAT('/cards/', CONVERT(mv.card_id, char)) AS url,
			(CASE WHEN ISNULL(ban.card_id) THEN 3 ELSE ban.status END) AS limitation,
			mv.deck_amount AS amount,
			mv.card_amount / mv.deck_amount AS average
		FROM mv_top_archetype_cards mv
		JOIN entity_cards ec ON ec.id = mv.card_id
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS ban ON ban.card_id = mv.card_id AND ban.banlist_id = ?
			WHERE mv.archetype_id = ? AND NOT match_archetype(ec.set_code, mv.archetype_id)
			ORDER BY mv.deck_amount DESC, mv.card_amount DESC, mv.card_id ASC;
	`, util.GodotEnv("LAST_BANLIST"), input.ID).Find(&topArchetypeCards)

	return &topArchetypeCards
}
