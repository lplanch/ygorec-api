package model

type EnumRule struct {
	ID    uint64 `gorm:"type:INTEGER PRIMARY KEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
