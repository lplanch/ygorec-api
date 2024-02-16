package model

type EntityCard struct {
	ID        uint32 `gorm:"type:INTEGER PRIMARYKEY"`
	Name      string `gorm:"type:TEXT;NOT NULL"`
	Desc      string `gorm:"type:TEXT;NOT NULL"`
	SetCode   uint32 `gorm:"type:INTEGER;NOT NULL"`
	Type      uint32 `gorm:"type:INTEGER;NOT NULL"`
	Atk       uint32 `gorm:"type:INTEGER;NOT NULL"`
	Def       uint32 `gorm:"type:INTEGER;NOT NULL"`
	Level     uint32 `gorm:"type:INTEGER;NOT NULL"`
	Race      uint32 `gorm:"type:INTEGER;NOT NULL"`
	Attribute uint32 `gorm:"type:INTEGER;NOT NULL"`
	Category  uint32 `gorm:"type:INTEGER;NOT NULL"`
}
