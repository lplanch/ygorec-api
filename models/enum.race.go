package model

type EnumRace struct {
	ID    uint64 `gorm:"type:INTEGER PRIMARY KEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
