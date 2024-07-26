package types

func IsRelativeSize(l Literal) bool {
	return l == "smaller" || l == "larger"
}
