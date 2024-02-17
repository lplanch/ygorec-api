package model

type ModelCard struct {
	ID         uint64
	Name       string
	Desc       string
	Attribute  string
	Types      []string
	Race       string
	Archetypes []string
	Atk        int64
	Def        int64
	Level      string
	Categories []string
}

type ModelDbCard struct {
	ID         uint64
	Name       string
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
	ID   uint64
	Name string
}
