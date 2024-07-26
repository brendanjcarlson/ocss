package types

func IsAnglePercentage(l Literal) bool {
	if IsAngle(l) {
		return true
	}
	if IsPercentage(l) {
		return true
	}
	return false
}
