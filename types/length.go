package types

import "strings"

func IsLength(l Literal) bool {
	str := l.String()
	for _, unit := range lengthUnits {
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

func IsLengthUnit(l Literal) bool {
	str := l.String()
	for _, unit := range lengthUnits {
		if str == unit {
			return true
		}
	}
	return false
}

var lengthUnits = [49]string{
	"cap",
	"ch",
	"em",
	"ex",
	"ic",
	"lh",
	"rcap",
	"rch",
	"rem",
	"rex",
	"ric",
	"rlh",
	"vh",
	"vw",
	"vmax",
	"vmin",
	"vb",
	"vi",
	"cqw",
	"cqh",
	"cqi",
	"cqb",
	"cqmin",
	"cqmax",
	"px",
	"cm",
	"mm",
	"Q",
	"in",
	"pc",
	"pt",
	"dvb",
	"dvh",
	"dvi",
	"dvmax",
	"dvmin",
	"dvw",
	"lvb",
	"lvh",
	"lvi",
	"lvmax",
	"lvmin",
	"lvw",
	"svb",
	"svh",
	"svi",
	"svmax",
	"svmin",
	"svw",
}
