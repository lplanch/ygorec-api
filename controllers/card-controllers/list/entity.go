package listCards

type InputListCards struct {
	Limit   int `validate:"gt=0"`
	Offset  int `validate:"gte=0"`
	Banlist string
}
