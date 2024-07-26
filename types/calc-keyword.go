package types

var calcKeywords = [5]Literal{"e", "pi", "infinity", "-infinity", "NaN"}

func IsCalcKeyword(l Literal) bool {
	for _, keyword := range calcKeywords {
		if l == keyword {
			return true
		}
	}
	return false
}
