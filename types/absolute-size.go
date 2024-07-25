package types

type AbsoluteSize string

func (a AbsoluteSize) IsAbsoluteSize() bool {
	_, ok := absoluteSizeEnumValidationMap[a]
	return ok
}

const (
	AbsoluteSizeXXSmall  AbsoluteSize = "xx-small"
	AbsoluteSizeXSmall   AbsoluteSize = "x-small"
	AbsoluteSizeSmall    AbsoluteSize = "small"
	AbsoluteSizeMedium   AbsoluteSize = "medium"
	AbsoluteSizeLarge    AbsoluteSize = "large"
	AbsoluteSizeXLarge   AbsoluteSize = "x-large"
	AbsoluteSizeXXLarge  AbsoluteSize = "xx-large"
	AbsoluteSizeXXXLarge AbsoluteSize = "xxx-large"
)

var absoluteSizeEnumValidationMap = map[AbsoluteSize]bool{
	AbsoluteSizeXXSmall:  true,
	AbsoluteSizeXSmall:   true,
	AbsoluteSizeSmall:    true,
	AbsoluteSizeMedium:   true,
	AbsoluteSizeLarge:    true,
	AbsoluteSizeXLarge:   true,
	AbsoluteSizeXXLarge:  true,
	AbsoluteSizeXXXLarge: true,
}

var absoluteSizePropertyValidationMap = map[Property]bool{
	PropertyFont:     true,
	PropertyFontSize: true,
}
