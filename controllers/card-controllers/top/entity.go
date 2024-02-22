package listCards

import "time"

type InputListCards struct {
	From   time.Time `validate:"datetime"`
	To     time.Time `validate:"datetime"`
	Limit  uint16    `validate:"gt=0"`
	Offset uint16    `validate:"gte=0"`
}
