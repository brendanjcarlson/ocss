package token

import "fmt"

type Token struct {
	kind     Kind
	literal  Literal
	position Position
}

func New(kind Kind, literal Literal, position Position) *Token {
	return &Token{kind, literal, position}
}

func (t *Token) Debug() string {
	return fmt.Sprintf("%s %s", t.String(), t.position.String())
}

func (t *Token) Kind() Kind {
	return t.kind
}

func (t *Token) Literal() Literal {
	return t.literal
}

func (t *Token) Position() Position {
	return t.position
}

func (t *Token) String() string {
	return fmt.Sprintf("[ KIND: %s | LITERAL: %q ]", t.kind.String(), t.Literal())
}

func (t *Token) Type() string {
	if t.kind == KIND_DIMENSION {
		return "has_type"
	} else {
		return "typeless"
	}
}
