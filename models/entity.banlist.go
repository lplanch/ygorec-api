package model

type EntityBanlist struct {
	ID string `gorm:"type:VARCHAR(255) PRIMARY KEY"`
	Ot uint64 `gorm:"type:BIGINT UNSIGNED"`
}
