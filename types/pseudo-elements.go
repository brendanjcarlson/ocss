package types

func IsPseudoElement(l Literal) bool {
	if _, ok := pseudoelementsValidationMap[l]; ok {
		return true
	}
	return false
}

var pseudoelementsValidationMap = map[Literal]bool{
	"-moz-color-swatch":                  true,
	"-moz-focus-inner":                   true,
	"-moz-list-bullet":                   true,
	"-moz-list-number":                   true,
	"-moz-progress-bar":                  true,
	"-moz-range-progress":                true,
	"-moz-range-thumb":                   true,
	"-moz-range-track":                   true,
	"-webkit-inner-spin-button":          true,
	"-webkit-meter-bar":                  true,
	"-webkit-meter-even-less-good-value": true,
	"-webkit-meter-inner-element":        true,
	"-webkit-meter-optimum-value":        true,
	"-webkit-meter-suboptimum-value":     true,
	"-webkit-progress-bar":               true,
	"-webkit-progress-inner-element":     true,
	"-webkit-progress-value":             true,
	"-webkit-scrollbar":                  true,
	"-webkit-search-cancel-button":       true,
	"-webkit-search-results-button":      true,
	"-webkit-slider-runnable-track":      true,
	"-webkit-slider-thumb":               true,
	"after":                              true,
	"backdrop":                           true,
	"before":                             true,
	"cue":                                true,
	"file-selector-button":               true,
	"first-letter":                       true,
	"first-line":                         true,
	"grammar-error":                      true,
	"highlight":                          true,
	"marker":                             true,
	"part":                               true,
	"placeholder":                        true,
	"selection":                          true,
	"slotted":                            true,
	"spelling-error":                     true,
	"target-text":                        true,
	"view-transition":                    true,
	"view-transition-group":              true,
	"view-transition-image-pair":         true,
	"view-transition-new":                true,
	"view-transition-old":                true,
}
