package model

type ModelCard struct {
	ID         uint32
	Name       string
	Desc       string
	Archetypes []string
	Type       string
	Atk        uint32
	Def        uint32
	Level      uint32
	Race       string
	Attribute  string
	Category   []string
}
