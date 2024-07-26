package types

type Literal string

// String implements fmt.Stringer.
func (l Literal) String() string {
	return string(l)
}

type Validator func(l Literal) bool

func IsOneOf(needle Literal, haystack ...Literal) bool {
	if len(haystack) == 0 {
		return false
	}
	for _, straw := range haystack {
		if needle == straw {
			return true
		}
	}
	return false
}
