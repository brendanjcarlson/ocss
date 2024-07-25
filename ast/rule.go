package ast

import (
	"strings"

	"github.com/brendanjcarlson/ocss/token"
)

var _ StyleRuleNode = (*StyleRule)(nil)

type StyleRule struct {
	Parent       RuleNode
	Selectors    []SelectorNode
	Declarations []DeclarationNode
	Children     []RuleNode
}

// CSS implements RuleNode.
func (s *StyleRule) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements RuleNode.
func (s *StyleRule) Literal() token.Literal {
	depth := 0
	for parent := s.ParentNode(); parent != nil; parent = parent.ParentNode() {
		depth++
	}

	out := new(strings.Builder)
	for _, selector := range s.Selectors {
		out.WriteString("<selector: \"")
		out.WriteString(selector.Literal().String())
		out.WriteString("\">\n")
	}
	for _, declaration := range s.Declarations {
		out.WriteString("\t<declaration: \"")
		out.WriteString(declaration.Literal().String())
		out.WriteString("\">\n")
	}
	out.WriteString("\n")
	for _, child := range s.Children {
		out.WriteString(child.Literal().String())
	}
	out.WriteString("\n")

	return token.Literal(out.String())
}

func (s *StyleRule) ParentNode() RuleNode {
	return s.Parent
}

func (s *StyleRule) SelectorNodes() []SelectorNode {
	return s.Selectors
}

// node implements RuleNode.
func (s *StyleRule) node() {}

// ruleNode implements RuleNode.
func (s *StyleRule) ruleNode() {}

// styleRuleNode implements StyleRuleNode.
func (s *StyleRule) styleRuleNode() {}

var _ AtRuleNode = (*MediaOrSupportsAtRule)(nil)

type MediaOrSupportsAtRule struct {
	Parent     RuleNode
	AtKeyword  *token.Token
	Condition  []*token.Token
	StyleRules []StyleRuleNode
}

// CSS implements AtRuleNode.
func (m *MediaOrSupportsAtRule) CSS(minified bool) string {
	panic("unimplemented")
}

// Literal implements AtRuleNode.
func (m *MediaOrSupportsAtRule) Literal() token.Literal {
	out := new(strings.Builder)
	out.WriteString("<at-keyword: \"")
	out.WriteString(m.AtKeyword.Literal().String())
	out.WriteString("\">")
	out.WriteString("<condition: \"")
	for i, condition := range m.Condition {
		out.WriteString(condition.Literal().String())
		if condition.Kind() == token.KIND_IDENTIFIER &&
			(i+1 < len(m.Condition)-1 && m.Condition[i+1].Kind() != token.KIND_COLON) {
			out.WriteString(" ")
		}
	}
	out.WriteString("\">\n")
	for _, style := range m.StyleRules {
		out.WriteString(style.Literal().String())
	}

	return token.Literal(out.String())
}

// ParentNode implements AtRuleNode.
func (m *MediaOrSupportsAtRule) ParentNode() RuleNode {
	return m.Parent
}

// atRuleNode implements AtRuleNode.
func (m *MediaOrSupportsAtRule) atRuleNode() {}

// node implements AtRuleNode.
func (m *MediaOrSupportsAtRule) node() {}

// ruleNode implements AtRuleNode.
func (m *MediaOrSupportsAtRule) ruleNode() {}
