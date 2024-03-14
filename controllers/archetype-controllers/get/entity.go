package getArchetype

type InputGetArchetype struct {
	Value  string `validate:"required"`
	Limit  int    `validate:"number,gt=0"`
	Offset int    `validate:"number,gte=0"`
}

type InputServiceGetArchetype struct {
	ID     uint64
	Limit  int
	Offset int
}
