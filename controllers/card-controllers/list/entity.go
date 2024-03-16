package listCards

type InputListCards struct {
	Limit   int    `validate:"number,gt=0"`
	Offset  int    `validate:"number,gte=0"`
	Banlist string `validate:"banlistdate"`
	CardID  int    `validate:"number,gte=0"`
}
