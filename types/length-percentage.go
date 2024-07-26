package types

func IsLengthPercentage(l Literal) bool {
	if IsLength(l) {
		return true
	}
	if IsPercentage(l) {
		return true
	}
	return false
}
