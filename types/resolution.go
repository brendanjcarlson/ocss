package types

import "strings"

var resolutionUnits = [4]string{"dpi", "dpcm", "dppx", "x"}

func IsResolution(l Literal) bool {
	str := l.String()
	for _, unit := range resolutionUnits {
		if rest, ok := strings.CutSuffix(str, unit); ok {
			if IsNumberPositive(Literal(rest)) {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
