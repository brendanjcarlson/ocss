package types

type Property string

func (p Property) IsProperty() bool {
	_, ok := propertyEnumValidationMap[p]
	return ok
}

const (
	PropertyFont     Property = "font"
	PropertyFontSize Property = "font-size"
)

var propertyEnumValidationMap = map[Property]bool{
	PropertyFont:     true,
	PropertyFontSize: true,
}
