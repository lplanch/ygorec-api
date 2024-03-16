package model

type ModelCard struct {
	ID         uint64
	Name       string
	Limitation uint8
	Desc       string
	Attribute  string
	Types      []string
	Race       string
	Archetypes []ModelArchetype
	Atk        int64
	Def        int64
	Level      string
	Categories []string
}

type ModelDbCard struct {
	ID         uint64
	Name       string
	Limitation uint8
	Desc       string
	Attribute  string
	Types      string
	Race       string
	Archetypes string
	Atk        int64
	Def        int64
	Level      string
	Categories string
}

type ModelListCard struct {
	ID    uint64
	Label string
	Url   string
}

type ModelListCardStats struct {
	ID         uint64
	Label      string
	Url        string
	Limitation uint8
	Amount     uint32
	Average    float32
}

type ModelFullListCardStats struct {
	DeckAmount uint32
	List       []ModelListCardStats
}
