package enums

type GenderEnumServices int

const (
	MALE GenderEnumServices = iota
	FEMALE
	ALL
)

func (g GenderEnumServices) String() string {
	return [...]string{"MALE", "FEMALE", "ALL"}[g]
}

func (g GenderEnumServices) IsValid() bool {
	switch g {
	case MALE, FEMALE, ALL:
		return true
	}

	return false
}
