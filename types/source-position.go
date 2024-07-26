package types

import "fmt"

type SourceLocation struct {
	row   int
	start int
	end   int
}

func NewSourceLocation(row, start, end int) SourceLocation {
	return SourceLocation{row, start, end}
}

func (p SourceLocation) End() int {
	return p.end
}

func (p SourceLocation) Row() int {
	return p.row
}

func (p SourceLocation) Span() (start, end int) {
	return p.start, p.end
}

func (p SourceLocation) Start() int {
	return p.start
}

// String implements fmt.Stringer.
func (p SourceLocation) String() string {
	return fmt.Sprintf("[ ROW: %d | SPAN: %d -- %d ]", p.row, p.start, p.end)
}
