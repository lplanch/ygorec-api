package getCard

type InputGetCard struct {
	ID string `validate:"required,number"`
}

type InputServiceGetCard struct {
	ID uint64
}
