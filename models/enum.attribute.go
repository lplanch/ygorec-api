package model

type EnumAttribute struct {
	ID    uint32 `gorm:"type:INTEGER PRIMARYKEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
