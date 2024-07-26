package types

import "strings"

func IsPosition(l Literal) bool {
	values := strings.Split(l.String(), " ")
	switch len(values) {
	case 1:
		if IsPositionKeyword(l) {
			return true
		}
		if IsLengthPercentage(l) {
			return true
		}
		return false
	case 2:
		one := Literal(values[0])
		two := Literal(values[1])
		if IsOneOf(one, "left", "center", "right") && IsOneOf(two, "top", "center", "bottom") {
			return true
		}
		if (IsLengthPercentage(one) || IsOneOf(one, "left", "center", "right")) &&
			(IsLengthPercentage(two) || IsOneOf(two, "top", "center", "bottom")) {
			return true
		}
		return false
	case 4:
		one := Literal(values[0])
		two := Literal(values[1])
		three := Literal(values[2])
		four := Literal(values[3])
		if IsOneOf(one, "left", "right") && IsLengthPercentage(two) &&
			IsOneOf(three, "top", "bottom") && IsLengthPercentage(four) {
			return true
		}
		if IsOneOf(one, "top", "bottom") && IsLengthPercentage(two) &&
			IsOneOf(three, "left", "right") && IsLengthPercentage(four) {
			return true
		}
		return false
	}
	return false
}

var positionKeywords = [5]string{"left", "center", "right", "top", "bottom"}

func IsPositionKeyword(l Literal) bool {
	for _, keyword := range positionKeywords {
		if l.String() == keyword {
			return true
		}
	}
	return false
}

var positionKeywordsHorizontal = [3]string{"left", "center", "right"}

func IsPositionValueHorizontal(l Literal) bool {
	str := l.String()
	for _, keyword := range positionKeywordsHorizontal {
		if str == keyword {
			return true
		}
	}
	return false
}

var positionKeywordsVertical = [3]string{"top", "center", "bottom"}

func IsPositionValueVertical(l Literal) bool {
	str := l.String()
	for _, keyword := range positionKeywordsVertical {
		if str == keyword {
			return true
		}
	}
	return false
}
