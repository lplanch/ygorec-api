package listArchetypes

type InputListArchetypes struct {
	Limit  int `validate:"number,gt=0"`
	Offset int `validate:"number,gte=0"`
}
