package types

func IsAbsoluteSize(l Literal) bool {
	if _, ok := absoluteSizeEnumValidationMap[l]; ok {
		return true
	}
	return false
}

var absoluteSizeEnumValidationMap = map[Literal]bool{
	"xx-small":  true,
	"x-small":   true,
	"small":     true,
	"medium":    true,
	"large":     true,
	"x-large":   true,
	"xx-large":  true,
	"xxx-large": true,
}
