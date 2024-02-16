package model

type EntityCard struct {
	ID        uint64 `gorm:"type:INTEGER PRIMARYKEY"`
	Ot        uint64 `gorm:"type:INTEGER"`
	Alias     uint64 `gorm:"type:INTEGER"`
	SetCode   uint64 `gorm:"type:INTEGER"`
	Type      uint64 `gorm:"type:INTEGER"`
	Atk       int64  `gorm:"type:INTEGER"`
	Def       int64  `gorm:"type:INTEGER"`
	Level     uint64 `gorm:"type:INTEGER"`
	Race      uint64 `gorm:"type:INTEGER"`
	Attribute uint64 `gorm:"type:INTEGER"`
	Category  uint64 `gorm:"type:INTEGER"`
	Name      string `gorm:"type:TEXT"`
	Desc      string `gorm:"type:TEXT"`
}
