package types

import "regexp"

var isPercentageRegexp *regexp.Regexp

func initIsPercentageRegexp() {
	if isPercentageRegexp == nil {
		isPercentageRegexp = regexp.MustCompile(`[+-]?\.?\d+\.?\d*%`)
	}
}

func (l Literal) IsPercentage() bool {
	return isPercentageRegexp.MatchString(l.String())
}
