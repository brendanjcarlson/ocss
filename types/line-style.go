package types

func IsLineStyle(l Literal) bool {
	if _, ok := lineStyleEnumValidationMap[l]; ok {
		return true
	}
	return false
}

var lineStyleEnumValidationMap = map[Literal]bool{
	"none":   true,
	"hidden": true,
	"dotted": true,
	"dashed": true,
	"solid":  true,
	"double": true,
	"groove": true,
	"ridge":  true,
	"inset":  true,
	"outset": true,
}
