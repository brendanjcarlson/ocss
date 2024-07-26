package types

import (
	"regexp"
)

const isPerctageRegexpPattern = `^[+-]?(\d+(\.\d*)?|\.\d+)%$`

var isPercentageRegexp *regexp.Regexp

func initIsPercentageRegexp() {
	if isPercentageRegexp == nil {
		isPercentageRegexp = regexp.MustCompile(isPerctageRegexpPattern)
	}
}

func IsPercentage(l Literal) bool {
	// if the literal is shorter than a single digit and a percent sign
	// the literal cannot be a percentage
	if len(l) < 2 {
		return false
	}
	if isPercentageRegexp.MatchString(l.String()) {
		return true
	}
	return false
}
