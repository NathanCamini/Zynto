package models

type GenderEnum int

const (
	MALE = iota
	FEMALE
)

func (g GenderEnum) String() string {
	return [...]string{"MALE", "FEMALE"}[g]
}