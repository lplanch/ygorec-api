package listCards

import "time"

type InputListCards struct {
	From   time.Time `validate:"datetime"`
	To     time.Time `validate:"datetime"`
	Limit  int       `validate:"gt=0"`
	Offset int       `validate:"gte=0"`
}
