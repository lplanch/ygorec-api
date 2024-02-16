package model

type EntityCard struct {
	ID        uint32 `gorm:"type:INTEGER PRIMARYKEY"`
	Ot        uint32 `gorm:"type:INTEGER"`
	Alias     uint32 `gorm:"type:INTEGER"`
	SetCode   uint32 `gorm:"type:INTEGER"`
	Type      uint32 `gorm:"type:INTEGER"`
	Atk       uint32 `gorm:"type:INTEGER"`
	Def       uint32 `gorm:"type:INTEGER"`
	Level     uint32 `gorm:"type:INTEGER"`
	Race      uint32 `gorm:"type:INTEGER"`
	Attribute uint32 `gorm:"type:INTEGER"`
	Category  uint32 `gorm:"type:INTEGER"`
	Name      string `gorm:"type:TEXT"`
	Desc      string `gorm:"type:TEXT"`
}
