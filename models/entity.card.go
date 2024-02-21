package model

type EntityCard struct {
	ID        uint64 `gorm:"type:BIGINT UNSIGNED PRIMARY KEY"`
	Ot        uint64 `gorm:"type:BIGINT UNSIGNED"`
	Alias     uint64 `gorm:"type:BIGINT UNSIGNED"`
	SetCode   uint64 `gorm:"type:BIGINT UNSIGNED"`
	Type      uint64 `gorm:"type:BIGINT UNSIGNED"`
	Atk       int64  `gorm:"type:BIGINT"`
	Def       int64  `gorm:"type:BIGINT"`
	Level     uint64 `gorm:"type:BIGINT UNSIGNED"`
	Race      uint64 `gorm:"type:BIGINT UNSIGNED"`
	Attribute uint64 `gorm:"type:BIGINT UNSIGNED"`
	Category  uint64 `gorm:"type:BIGINT UNSIGNED"`
	Name      string `gorm:"type:TEXT"`
	Desc      string `gorm:"type:TEXT"`
}
