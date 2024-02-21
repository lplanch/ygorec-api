package model

type EnumType struct {
	ID    uint64 `gorm:"type:BIGINT UNSIGNED PRIMARY KEY"`
	Value string `gorm:"type:TEXT;NOT NULL"`
}
