package types

var overflowKeywords = [5]Literal{"visible", "hidden", "clip", "scroll", "auto"}

func IsOverflow(l Literal) bool {
	for _, keyword := range overflowKeywords {
		if l == keyword {
			return true
		}
	}
	return false
}
