package types

import "strings"

var timeUnits = [2]string{"s", "ms"}

func IsTime(l Literal) bool {
	str := l.String()
	for _, unit := range timeUnits {
		if rest, ok := strings.CutSuffix(str, unit); ok {
			if IsNumber(Literal(rest)) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
