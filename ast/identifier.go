package ast

import "github.com/brendanjcarlson/ocss/token"

var _ ExpressionNode = (*Identifier)(nil)

type Identifier struct {
	Token token.Token
}

// CSS implements ExpressionNode.
func (i *Identifier) CSS(minified bool) string {
	return i.Token.Literal().String()
}

// Literal implements ExpressionNode.
func (i *Identifier) Literal() token.Literal {
	return i.Token.Literal()
}

// expressionNode implements ExpressionNode.
func (i *Identifier) expressionNode() {}

// node implements ExpressionNode.
func (i *Identifier) node() {}
