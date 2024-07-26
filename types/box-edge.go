package types

func IsBoxEdge(l Literal) bool {
	if IsBoxEdgeVisualBoxEdge(l) {
		return true
	}
	if IsBoxEdgeLayoutBoxEdge(l) {
		return true
	}
	if IsBoxEdgePaintBoxEdge(l) {
		return true
	}
	if IsBoxEdgeCoordBoxEdge(l) {
		return true
	}
	if IsBoxEdgeGeometryBoxEdge(l) {
		return true
	}
	return false
}

var boxEdgesVisualBoxEdges = [3]Literal{"content-box", "padding-box", "border-box"}

func IsBoxEdgeVisualBoxEdge(l Literal) bool {
	for _, box := range boxEdgesVisualBoxEdges {
		if l == box {
			return true
		}
	}
	return false
}

var boxEdgesLayoutBoxEdges = [4]Literal{"context-box", "padding-box", "border-box", "margin-box"}

func IsBoxEdgeLayoutBoxEdge(l Literal) bool {
	for _, box := range boxEdgesLayoutBoxEdges {
		if l == box {
			return true
		}
	}
	return false
}

var boxEdgesPaintBoxEdges = [5]Literal{
	"context-box",
	"padding-box",
	"border-box",
	"fill-box",
	"stroke-box",
}

func IsBoxEdgePaintBoxEdge(l Literal) bool {
	for _, box := range boxEdgesPaintBoxEdges {
		if l == box {
			return true
		}
	}
	return false
}

var boxEdgesCoordBoxEdges = [6]Literal{
	"context-box",
	"padding-box",
	"border-box",
	"fill-box",
	"stroke-box",
	"view-box",
}

func IsBoxEdgeCoordBoxEdge(l Literal) bool {
	for _, box := range boxEdgesCoordBoxEdges {
		if l == box {
			return true
		}
	}
	return false
}

var boxEdgesGeometryBoxEdges = [3]Literal{"fill-box", "stroke-box", "view-box"}

func IsBoxEdgeGeometryBoxEdge(l Literal) bool {
	for _, edge := range boxEdgesGeometryBoxEdges {
		if l == edge {
			return true
		}
	}
	return false
}
