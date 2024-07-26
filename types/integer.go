package types

import "regexp"

const isIntegerRegexpPattern = `[+-]?\d+`

var isIntegerRegexp *regexp.Regexp

func initIsIntegerRegexp() {
	if isIntegerRegexp == nil {
		isIntegerRegexp = regexp.MustCompile(isIntegerRegexpPattern)
	}
}

func IsInteger(l Literal) bool {
	if isIntegerRegexp.MatchString(l.String()) {
		return true
	}
	return false
}
