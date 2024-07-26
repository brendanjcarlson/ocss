package types

func IsBlendMode(l Literal) bool {
	if _, ok := blendModeValidationMap[l]; ok {
		return true
	}
	return false
}

var blendModeValidationMap = map[Literal]bool{
	"normal":      true,
	"multiply":    true,
	"screen":      true,
	"overlay":     true,
	"darken":      true,
	"lighten":     true,
	"color-dodge": true,
	"color-burn":  true,
	"hard-light":  true,
	"soft-light":  true,
	"difference":  true,
	"exclusion":   true,
	"hue":         true,
	"saturation":  true,
	"color":       true,
	"luminosity":  true,
}
