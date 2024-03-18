package model

type ModelArchetype struct {
	ArchetypeID    uint64
	Label          string
	DeckAmount     uint32
	CardAmount     uint32
	MostUsedCardID uint64
	Url            string
}

type ModelFullListArchetypeCardStats struct {
	Label          string
	DeckAmount     uint32
	ArchetypeCards []ModelListCardStats
	OtherCards     []ModelListCardStats
}

type ModelListArchetype struct {
	ID    uint64
	Label string
	Url   string
}
