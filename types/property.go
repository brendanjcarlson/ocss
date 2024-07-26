package types

func IsProperty(l Literal) bool {
	if _, ok := propertyEnumValidationMap[l]; ok {
		return true
	}
	return false
}

var propertyEnumValidationMap = map[Literal]bool{
	"font":      true,
	"font-size": true,
}
