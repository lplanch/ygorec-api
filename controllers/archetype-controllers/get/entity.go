package getArchetype

type InputGetArchetype struct {
	Value string `validate:"required"`
}

type InputServiceGetArchetype struct {
	ID uint64
}
