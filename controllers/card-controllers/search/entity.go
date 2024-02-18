package searchCard

type InputSearchCard struct {
	Q string `validate:"required,lowercase"`
}
