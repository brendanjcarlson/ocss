package types

func IsSystemColor(l Literal) bool {
	if _, ok := systemColorEnumValidationMap[l]; ok {
		return true
	}
	return false
}

var systemColorEnumValidationMap = map[Literal]bool{
	"AccentColor":      true,
	"AccentColorText":  true,
	"ActiveText":       true,
	"ButtonBorder":     true,
	"ButtonFace":       true,
	"ButtonText":       true,
	"Canvas":           true,
	"CanvasText":       true,
	"Field":            true,
	"FieldText":        true,
	"GrayText":         true,
	"Highlight":        true,
	"HighlightText":    true,
	"LinkText":         true,
	"Mark":             true,
	"MarkText":         true,
	"SelectedItem":     true,
	"SelectedItemText": true,
	"VisitedText":      true,
}
