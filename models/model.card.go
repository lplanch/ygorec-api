package model

type ModelCard struct {
	ID         uint64
	Name       string
	Desc       string
	Archetypes []string
	Type       string
	Atk        uint64
	Def        uint64
	Level      uint64
	Race       string
	Attribute  string
	Category   []string
}
