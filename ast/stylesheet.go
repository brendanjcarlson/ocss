package ast

import (
	"strings"

	"github.com/brendanjcarlson/ocss/token"
)

var _ StyleSheetNode = (*StyleSheet)(nil)

type StyleSheet struct {
	Nodes []Node
}

// CSS implements StyleSheetNode.
func (s *StyleSheet) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements StyleSheetNode.
func (s *StyleSheet) Literal() token.Literal {
	out := new(strings.Builder)
	for _, node := range s.Nodes {
		out.WriteString(node.Literal().String())
	}
	return token.Literal(out.String())
}

// node implements StyleSheetNode.
func (s *StyleSheet) node() {}

// styleSheetNode implements StyleSheetNode.
func (s *StyleSheet) styleSheetNode() {}
