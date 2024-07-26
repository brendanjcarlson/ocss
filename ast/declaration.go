package ast

import (
	"strings"

	"github.com/brendanjcarlson/ocss/token"
	"github.com/brendanjcarlson/ocss/types"
)

// Enforce interface implementation at compile time.
var _ DeclarationNode = (*Declaration)(nil)

type Declaration struct {
	Property  *token.Token
	Colon     *token.Token
	Values    []*token.Token
	Semicolon *token.Token
}

// CSS implements DeclarationNode.
func (d *Declaration) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements DeclarationNode.
func (d *Declaration) Literal() types.Literal {
	out := new(strings.Builder)
	out.WriteString(d.Property.Literal().String())
	out.WriteString(d.Colon.Literal().String())
	for i, value := range d.Values {
		if i > 0 && i < len(d.Values) {
			out.WriteString(" ")
		}
		out.WriteString(value.Literal().String())
	}
	out.WriteString(d.Semicolon.Literal().String())
	return types.Literal(out.String())
}

// declarationNode implements DeclarationNode.
func (d *Declaration) declarationNode() {}

// node implements DeclarationNode.
func (d *Declaration) node() {}
