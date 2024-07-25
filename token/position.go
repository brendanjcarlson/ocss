package token

import "fmt"

type Position struct {
	row   int
	start int
	end   int
}

func NewPosition(row, start, end int) Position {
	return Position{row, start, end}
}

func (p Position) End() int {
	return p.end
}

func (p Position) Row() int {
	return p.row
}

func (p Position) Span() (start, end int) {
	return p.start, p.end
}

func (p Position) Start() int {
	return p.start
}

// String implements fmt.Stringer.
func (p Position) String() string {
	return fmt.Sprintf("[ ROW: %d | SPAN: %d -- %d ]", p.row, p.start, p.end)
}
