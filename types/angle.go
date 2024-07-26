package types

import (
	"regexp"
)

const isAngleRegexpPattern = `^[+-]?(\d+(\.\d*)?|\.\d+)(deg|rad|grad|turn)$`

var isAngleRegexp *regexp.Regexp

func initIsAngleRegexp() {
	if isAngleRegexp == nil {
		isAngleRegexp = regexp.MustCompile(isAngleRegexpPattern)
	}
}

func IsAngle(l Literal) bool {
	// if the literal is shorter than a single digit + the shortest unit
	// the literal cannot be an angle
	if len(l) < 4 {
		return false
	}
	if isAngleRegexp.MatchString(l.String()) {
		return true
	}
	return false
}
