package types

import "strings"

func IsRatio(l Literal) bool {
	if IsNumberPositive(l) {
		return true
	}
	if left, right, ok := strings.Cut(l.String(), "/"); ok {
		if IsNumberPositive(Literal(strings.TrimSpace(left))) &&
			IsNumberPositive(Literal(strings.TrimSpace(right))) {
			return true
		}
		return false
	}
	return false
}
