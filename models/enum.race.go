package model

type EnumRace struct {
	ID    uint32 `gorm:"type:INTEGER PRIMARYKEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
