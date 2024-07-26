package types

func IsAtKeyword(l Literal) bool {
	if _, ok := atKeywordsValidationMap[l]; ok {
		return true
	}
	return false
}

var atKeywordsValidationMap = map[Literal]bool{
	"charset":             true,
	"color-profile":       true,
	"container":           true,
	"counter-style":       true,
	"document":            true,
	"font-face":           true,
	"font-feature-values": true,
	"swash":               true,
	"annotation":          true,
	"ornaments":           true,
	"stylistic":           true,
	"styleset":            true,
	"character-variant":   true,
	"font-palette-values": true,
	"import":              true,
	"keyframes":           true,
	"layer":               true,
	"media":               true,
	"namespace":           true,
	"page":                true,
	"position-try":        true,
	"property":            true,
	"scope":               true,
	"starting-style":      true,
	"supports":            true,
	"view-transition":     true,
}
