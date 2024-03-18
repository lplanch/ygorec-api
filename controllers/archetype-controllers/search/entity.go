package searchArchetype

type InputSearchArchetype struct {
	Q string `validate:"required,lowercase"`
}
