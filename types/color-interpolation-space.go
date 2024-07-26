package types

var colorInterpolationMethodRectangularColorSpaces = [11]Literal{
	"srgb",
	"srgb-linear",
	"display-p3",
	"a98-rgb",
	"prophoto-rgb",
	"rec2020",
	"lab",
	"oklab",
	"xyz",
	"xyz-d50",
	"xyz-d65",
}

func IsColorInterpolationMethodRectangularColorSpace(l Literal) bool {
	for _, space := range colorInterpolationMethodRectangularColorSpaces {
		if l == space {
			return true
		}
	}
	return false
}

var colorInterpolationMethodPolarColorSpaces = [4]Literal{"hsl", "hwb", "lch", "oklch"}

func IsColorInterpolationMethodPolarColorSpace(l Literal) bool {
	for _, space := range colorInterpolationMethodPolarColorSpaces {
		if l == space {
			return true
		}
	}
	return false
}

var colorInterpolationMethodHueInterpolationMethods = [5]Literal{
	"shorter",
	"longer",
	"increasing",
	"decreasing",
	"hue",
}

func IsColorInterpolationMethodHueInterpolationMethod(l Literal) bool {
	for _, method := range colorInterpolationMethodHueInterpolationMethods {
		if l == method {
			return true
		}
	}
	return false
}
