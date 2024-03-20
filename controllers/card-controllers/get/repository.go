package getCard

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetCardRepository(input *model.EntityCard) (*model.EntityCard, string)
	SanitizeCardRepository(input *model.EntityCard) (*model.ModelDbCard, string)
	GetListCardAlias(alias_ids []string) *[]model.ModelListCard
	GetCardArchetypes(archetype_names []string) *[]model.ModelArchetype
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryGet(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetCardRepository(input *model.EntityCard) (*model.EntityCard, string) {

	var card model.EntityCard

	db := r.db.Model(&card)
	errorCode := make(chan string, 1)

	getCard := db.Debug().Where("id = ?", input.ID).Find(&card)

	if getCard.RowsAffected < 1 {
		errorCode <- "GET_CARD_NOT_FOUND_404"
		return &card, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &card, <-errorCode
}

func (r *repository) SanitizeCardRepository(input *model.EntityCard) (*model.ModelDbCard, string) {

	var card model.ModelDbCard

	db := r.db
	errorCode := make(chan string, 1)

	getCard := db.Debug().Model(&input).Select(`
		entity_cards.id,
		entity_cards.name,
		(
			CASE WHEN ISNULL(p_ec.id)
			THEN (SELECT GROUP_CONCAT(ec.id) FROM entity_cards ec WHERE ec.alias = entity_cards.id)
			ELSE (SELECT GROUP_CONCAT(ec.id) FROM entity_cards ec WHERE entity_cards.id != ec.id AND (p_ec.id = ec.id OR ec.alias = p_ec.id))
			END
		) AS alias,
		(CASE WHEN ISNULL(b.card_id) THEN 3 ELSE b.status END) AS limitation,
		entity_cards.desc,
		(SELECT enum_attributes.value FROM enum_attributes WHERE entity_cards.attribute = enum_attributes.id LIMIT 1) AS attribute,
		(SELECT GROUP_CONCAT(enum_types.value) FROM enum_types WHERE entity_cards.type & enum_types.id) AS types,
		(SELECT enum_races.value FROM enum_races WHERE entity_cards.race = enum_races.id LIMIT 1) AS race,
		(SELECT GROUP_CONCAT(enum_archetypes.value) FROM enum_archetypes WHERE match_archetype(entity_cards.set_code, enum_archetypes.id)) AS archetypes,
		entity_cards.atk,
		entity_cards.def,
		(SELECT enum_levels.value FROM enum_levels WHERE entity_cards.level = enum_levels.id LIMIT 1) AS level,
		(SELECT GROUP_CONCAT(enum_categories.value) FROM enum_categories WHERE entity_cards.category & enum_categories.id) AS categories
	`).Joins(`
		LEFT OUTER JOIN entity_cards AS p_ec ON p_ec.id = entity_cards.alias
	`).Joins(`
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS b ON b.card_id = entity_cards.id AND b.banlist_id = ?
	`, util.GodotEnv("LAST_BANLIST")).Where("entity_cards.id = ?", input.ID).Find(&card)

	if getCard.RowsAffected < 1 {
		errorCode <- "GET_CARD_NOT_FOUND_404"
		return &card, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &card, <-errorCode
}

func (r *repository) GetListCardAlias(alias_ids []string) *[]model.ModelListCard {
	var cards []model.ModelListCard

	db := r.db.Model(&model.EntityCard{})

	db.Debug().Select(`
		id,
		name AS label,
		CONCAT('/cards/', CONVERT(id, char)) AS url
	`).Where(`
		id IN ?
	`, alias_ids).Order(`
		alias
	`).Find(&cards)

	return &cards
}

func (r *repository) GetCardArchetypes(archetype_names []string) *[]model.ModelArchetype {

	var archetypes []model.ModelArchetype

	db := r.db.Model(&model.MvTopArchetype{})

	db.Debug().Select(`
		archetype_id,
		a.value AS label,
		deck_amount,
		card_amount,
		most_used_card_id,
		CONCAT('/archetypes/', CONVERT(LOWER(REPLACE(a.value, ' ', '-')), char)) AS url
	`).Joins(`
		JOIN enum_archetypes a ON a.value IN ?
	`, archetype_names).Where(`
		archetype_id = a.id
	`).Find(&archetypes)

	return &archetypes
}
