package enums

type GenderEnum int

const (
	MALE GenderEnum = iota
	FEMALE
)

func (g GenderEnum) String() string {
	return [...]string{"MALE", "FEMALE"}[g]
}

func (g GenderEnum) IsValid() bool {
	switch g {
	case MALE, FEMALE:
		return true
	}

	return false
}
