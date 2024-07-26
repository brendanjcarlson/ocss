package types

import (
	"regexp"
)

const (
	isNumberRegexpPattern         = `^[+-]?(\d+(\.\d*)?|\.\d+)$`
	isNumberNegativeRegexpPattern = `^-(\d+(\.\d*)?|\.\d+)$`
	isNumberPositiveRegexpPattern = `^+?(\d+(\.\d*)?|\.\d+)$`
)

var (
	isNumberRegexp         *regexp.Regexp
	isNumberNegativeRegexp *regexp.Regexp
	isNumberPositiveRegexp *regexp.Regexp
)

func initIsNumberRegexp() {
	if isNumberRegexp == nil {
		isNumberRegexp = regexp.MustCompile(isNumberRegexpPattern)
	}
	if isNumberNegativeRegexp == nil {
		isNumberNegativeRegexp = regexp.MustCompile(isNumberNegativeRegexpPattern)
	}
	if isNumberPositiveRegexp == nil {
		isNumberPositiveRegexp = regexp.MustCompile(isNumberPositiveRegexpPattern)
	}
}

func IsNumber(l Literal) bool {
	if len(l) == 0 {
		return false
	}
	if isNumberRegexp.MatchString(l.String()) {
		return true
	}
	return false
}

func IsNumberNegative(l Literal) bool {
	if len(l) == 0 {
		return false
	}
	if isNumberNegativeRegexp.MatchString(l.String()) {
		return true
	}
	return false
}

func IsNumberPositive(l Literal) bool {
	if len(l) == 0 {
		return false
	}
	if isNumberPositiveRegexp.MatchString(l.String()) {
		return true
	}
	return false
}
