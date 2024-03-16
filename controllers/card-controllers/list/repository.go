package listCards

import (
	model "github.com/lplanch/test-go-api/models"
	util "github.com/lplanch/test-go-api/utils"
	"gorm.io/gorm"
)

type Repository interface {
	GetDeckAmount(input *InputListCards) *uint32
	GetCardDeckAmount(input *InputListCards) *uint32
	ListCardsRepository(input *InputListCards) *[]model.ModelListCardStats
	ListRelatedCardsRepository(input *InputListCards) *[]model.ModelListCardStats
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryList(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetDeckAmount(input *InputListCards) *uint32 {

	var total uint32

	db := r.db.Model(&model.EntityDeck{})

	db.Debug().Select(`
		COUNT(*)
	`).Where(`
		"" = ? OR updated_at > ?
	`, input.Banlist, input.Banlist).Find(&total)

	return &total
}

func (r *repository) GetCardDeckAmount(input *InputListCards) *uint32 {

	var total uint32

	db := r.db.Model(&model.GraphCardsBelongToDecks{})

	db.Debug().Select(`
		COUNT(DISTINCT deck_id)
	`).Where(`
		card_id = ?
	`, input.CardID).Find(&total)

	return &total
}

func (r *repository) ListCardsRepository(input *InputListCards) *[]model.ModelListCardStats {

	var cards []model.ModelListCardStats

	db := r.db.Model(&model.MvTopCard{})

	db.Debug().Select(`
		mv_top_cards.card_id AS id,
		e.name AS label,
		CONCAT('/cards/', CONVERT(mv_top_cards.card_id, char)) AS url,
		(CASE WHEN ISNULL(b.card_id) THEN 3 ELSE b.status END) AS limitation,
		mv_top_cards.amount,
		mv_top_cards.average
	`).Joins(`
		JOIN entity_cards e ON e.id = mv_top_cards.card_id
	`).Joins(`
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS b ON b.card_id = mv_top_cards.card_id AND b.banlist_id = ?
	`, util.GodotEnv("LAST_BANLIST")).Where(`
		IFNULL(mv_top_cards.banlist_id, "") = ?
	`, input.Banlist).Order(`
		mv_top_cards.amount DESC,
		mv_top_cards.card_id ASC
	`).Limit(input.Limit).Offset(input.Offset).Find(&cards)

	return &cards
}

func (r *repository) ListRelatedCardsRepository(input *InputListCards) *[]model.ModelListCardStats {

	var cards []model.ModelListCardStats

	db := r.db.Model(&model.MvTopRelatedCards{})

	db.Debug().Select(`
		to_card_id AS id,
		e.name AS label,
		CONCAT('/cards/', CONVERT(to_card_id, char)) AS url,
		(CASE WHEN ISNULL(ban.card_id) THEN 3 ELSE ban.status END) AS limitation,
		deck_amount AS amount,
		card_amount / deck_amount AS average
	`).Joins(`
		JOIN entity_cards e ON e.id = to_card_id
	`).Joins(`
		LEFT OUTER JOIN graph_cards_belong_to_banlists AS ban ON ban.card_id = to_card_id AND ban.banlist_id = ?
	`, util.GodotEnv("LAST_BANLIST")).Where(`
		from_card_id = ?
	`, input.CardID).Order(`
		amount DESC
	`).Order(`
		average DESC
	`).Order(`
		to_card_id ASC
	`).Limit(input.Limit).Offset(input.Offset).Find(&cards)

	return &cards
}
