package token

import (
	"fmt"

	"github.com/brendanjcarlson/ocss/types"
)

type Token struct {
	kind           Kind
	literal        types.Literal
	sourceLocation types.SourceLocation
}

func New(kind Kind, literal types.Literal, sourceLocation types.SourceLocation) *Token {
	return &Token{kind, literal, sourceLocation}
}

func (t *Token) Debug() string {
	return fmt.Sprintf("%s %s", t.String(), t.sourceLocation.String())
}

func (t *Token) Kind() Kind {
	return t.kind
}

func (t *Token) Literal() types.Literal {
	return t.literal
}

func (t *Token) SourceLocation() types.SourceLocation {
	return t.sourceLocation
}

func (t *Token) String() string {
	return fmt.Sprintf("[ KIND: %s | LITERAL: %q ]", t.kind.String(), t.Literal())
}
