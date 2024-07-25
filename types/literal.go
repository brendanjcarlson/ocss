package types

type Literal string

// String implements fmt.Stringer.
func (l Literal) String() string {
	return string(l)
}
