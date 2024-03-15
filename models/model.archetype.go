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
	DeckAmount     uint32
	ArchetypeCards []ModelListCardStats
	OtherCards     []ModelListCardStats
}
