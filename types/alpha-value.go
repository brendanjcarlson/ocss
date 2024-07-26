package types

func IsAlphaValue(l Literal) bool {
	if IsNumber(l) {
		return true
	}
	if IsPercentage(l) {
		return true
	}
	return false
}
