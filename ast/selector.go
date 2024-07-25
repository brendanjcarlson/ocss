package ast

import (
	"strings"

	"github.com/brendanjcarlson/ocss/token"
)

var _ SelectorNode = (*ElementSelector)(nil)

type ElementSelector struct {
	Parent StyleRuleNode
	Token  *token.Token
}

// CSS implements SelectorNode.
func (e *ElementSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (e *ElementSelector) Literal() token.Literal {
	return e.Token.Literal()
}

// ParentNode implements SelectorNode.
func (e *ElementSelector) ParentNode() StyleRuleNode {
	return e.Parent
}

// node implements SelectorNode.
func (e *ElementSelector) node() {}

// selectorNode implements SelectorNode.
func (e *ElementSelector) selectorNode() {}

var _ SelectorNode = (*ClassSelector)(nil)

type ClassSelector struct {
	Parent     StyleRuleNode
	Identifier *token.Token
}

// CSS implements SelectorNode.
func (c *ClassSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (c *ClassSelector) Literal() token.Literal {
	return c.Identifier.Literal()
}

// ParentNode implements SelectorNode.
func (c *ClassSelector) ParentNode() StyleRuleNode {
	return c.Parent
}

// node implements SelectorNode.
func (c *ClassSelector) node() {}

// selectorNode implements SelectorNode.
func (c *ClassSelector) selectorNode() {}

var _ SelectorNode = (*IdSelector)(nil)

type IdSelector struct {
	Parent     StyleRuleNode
	Hash       *token.Token
	Identifier *token.Token
}

// CSS implements SelectorNode.
func (i *IdSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (i *IdSelector) Literal() token.Literal {
	return i.Hash.Literal() + i.Identifier.Literal()
}

// ParentNode implements SelectorNode.
func (i *IdSelector) ParentNode() StyleRuleNode {
	return i.Parent
}

// node implements SelectorNode.
func (i *IdSelector) node() {}

// selectorNode implements SelectorNode.
func (i *IdSelector) selectorNode() {}

var _ SelectorNode = (*PseudoClassSelector)(nil)

type PseudoClassSelector struct {
	Parent     StyleRuleNode
	Colon      *token.Token
	Identifier *token.Token
}

// CSS implements SelectorNode.
func (p *PseudoClassSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (p *PseudoClassSelector) Literal() token.Literal {
	return p.Colon.Literal() + p.Identifier.Literal()
}

// ParentNode implements SelectorNode.
func (p *PseudoClassSelector) ParentNode() StyleRuleNode {
	return p.Parent
}

// node implements SelectorNode.
func (p *PseudoClassSelector) node() {}

// selectorNode implements SelectorNode.
func (p *PseudoClassSelector) selectorNode() {}

var _ SelectorNode = (*PseudoElementSelector)(nil)

type PseudoElementSelector struct {
	Parent     StyleRuleNode
	ColonOne   *token.Token
	ColonTwo   *token.Token
	Identifier *token.Token
}

// CSS implements SelectorNode.
func (p *PseudoElementSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (p *PseudoElementSelector) Literal() token.Literal {
	return p.ColonOne.Literal() + p.ColonTwo.Literal() + p.Identifier.Literal()
}

// ParentNode implements SelectorNode.
func (p *PseudoElementSelector) ParentNode() StyleRuleNode {
	return p.Parent
}

// node implements SelectorNode.
func (p *PseudoElementSelector) node() {}

// selectorNode implements SelectorNode.
func (p *PseudoElementSelector) selectorNode() {}

var _ SelectorNode = (*CombinatorSelector)(nil)

type CombinatorSelector struct {
	Parent StyleRuleNode
	Tokens []*token.Token
}

// CSS implements SelectorNode.
func (c *CombinatorSelector) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements SelectorNode.
func (c *CombinatorSelector) Literal() token.Literal {
	out := new(strings.Builder)
	for _, tok := range c.Tokens {
		out.WriteString(tok.Literal().String())
	}
	return token.Literal(out.String())
}

// ParentNode implements SelectorNode.
func (c *CombinatorSelector) ParentNode() StyleRuleNode {
	return c.Parent
}

// node implements SelectorNode.
func (c *CombinatorSelector) node() {}

// selectorNode implements SelectorNode.
func (c *CombinatorSelector) selectorNode() {}
