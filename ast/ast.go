package ast

import "github.com/brendanjcarlson/ocss/types"

type Node interface {
	CSS(minified bool) string
	Literal() types.Literal
	node()
}

type StyleSheetNode interface {
	Node
	styleSheetNode()
}

type RuleNode interface {
	Node
	ParentNode() RuleNode
	ruleNode()
}

type StyleRuleNode interface {
	RuleNode
	SelectorNodes() []SelectorNode
	styleRuleNode()
}

type AtRuleNode interface {
	RuleNode
	atRuleNode()
}

type SelectorNode interface {
	Node
	ParentNode() StyleRuleNode
	selectorNode()
}

type DeclarationNode interface {
	Node
	declarationNode()
}

type ExpressionNode interface {
	Node
	expressionNode()
}
