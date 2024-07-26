package parser

import (
	"fmt"

	"github.com/brendanjcarlson/ocss/ast"
	"github.com/brendanjcarlson/ocss/token"
	"github.com/brendanjcarlson/ocss/types"
)

type Parser struct {
	pos        int
	cursor     int
	tokens     []*token.Token
	tok        *token.Token
	stylesheet *ast.StyleSheet
}

func New(tokens []*token.Token) *Parser {
	p := &Parser{0, 0, tokens, nil, &ast.StyleSheet{}}
	p.consume()
	return p
}

func (p *Parser) ParseStyleSheet() *ast.StyleSheet {
main:
	for {
		p.skipWhitespace()

		if p.atEOF() {
			break main
		}

		if p.tok.Kind() == token.KIND_AT_KEYWORD {
			p.emit(p.consumeAtRule(nil))
		} else {
			p.emit(p.consumeStyleRule(nil))
		}
	}

	return p.stylesheet
}

func (p *Parser) consume() {
	if len(p.tokens) == 0 {
		p.tok = token.New(token.KIND_EOF, types.Literal(""), types.NewSourceLocation(0, 0, 0))
	}
	if p.cursor >= len(p.tokens) {
		p.tok = token.New(token.KIND_EOF, types.Literal(""), types.NewSourceLocation(0, 0, 0))
	} else {
		p.tok = p.tokens[p.cursor]
	}
	p.pos = p.cursor
	p.cursor += 1

	// NOTE: DEBUG
	fmt.Println(p.tok)
}

func (p *Parser) consumeStyleRule(parent ast.StyleRuleNode) *ast.StyleRule {
	// NOTE: DEBUG
	fmt.Println("consuming style rule")

	s := &ast.StyleRule{
		Parent: parent,
	}

	p.consumeSelectors(s)

	if !p.tokIs(token.KIND_CURLY_BRACKET_OPEN) {
		panic(p.unexpectedToken(p.tok, ", followed by selector or {"))
	}
	p.consume() // consume {

	p.consumeDeclarations(s)

	// TODO: handle nesting
	if !p.tokIs(token.KIND_CURLY_BRACKET_CLOSE) {
		panic(fmt.Sprintf("consumeStyleRule: expected }, got %s", p.tok))
	}

	p.consume()

	return s
}

func (p *Parser) consumeSelectors(parent *ast.StyleRule) {
	// NOTE: DEBUG
	fmt.Println("consuming selectors")

	for !p.atEOF() || p.tokIs(token.KIND_COMMA) {
		if p.tokIs(token.KIND_COMMA) {
			p.consume() // consume , before consuming next selector
		}
		selector := p.consumeSelector(parent)
		if selector == nil {
			return
		} else {
			parent.Selectors = append(parent.Selectors, selector)
		}
	}
}

func (p *Parser) consumeSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming selector")

	switch {
	case p.tokIs(token.KIND_IDENTIFIER) && types.IsHtmlElement(p.tok.Literal()):
		return p.consumeElementSelector(parent)
	case len(p.tok.Literal()) > 0 && p.tok.Literal().String()[0] == '.':
		return p.consumeClassSelector(parent)
	case p.tokIs(token.KIND_HASH):
		return p.consumeIdSelector(parent)
	case p.tokIs(token.KIND_COLON) && !p.peekIs(token.KIND_COLON):
		return p.consumePseudoClassSelector(parent)
	case p.tokIs(token.KIND_COLON) && p.peekIs(token.KIND_COLON):
		return p.consumePseudoElementSelector(parent)
		// TODO: ATTRIBUTE SELECTORS
	default:
		return nil
	}
}

func (p *Parser) consumeElementSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming element selector")

	if p.peekIs(token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		if !types.IsHtmlElement(p.tok.Literal()) {
			panic(p.unexpectedToken(p.tok, "html element"))
		}
		tok := p.tok
		p.consume()
		return &ast.ElementSelector{Parent: parent, Token: tok}
	} else {
		return p.consumeCombinatorSelector(parent)
	}
}

func (p *Parser) consumeClassSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming class selector")

	if p.tokIs(token.KIND_IDENTIFIER) &&
		p.peekIs(token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		identifier := p.tok
		p.consume() // advance to next token
		return &ast.ClassSelector{Parent: parent, Identifier: identifier}
	} else {
		return p.consumeCombinatorSelector(parent)
	}
}

func (p *Parser) consumeIdSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming id selector")

	if p.peekIs(token.KIND_IDENTIFIER) &&
		p.peekNIs(2, token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		hash := p.tok
		p.consume() // consume #
		identifier := p.tok
		p.consume() // consume identifier
		return &ast.IdSelector{Parent: parent, Hash: hash, Identifier: identifier}
	} else {
		return p.consumeCombinatorSelector(parent)
	}
}

func (p *Parser) consumePseudoClassSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming pseudo class selector")

	if p.peekIs(token.KIND_IDENTIFIER) &&
		p.peekNIs(2, token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		colon := p.tok
		p.consume()
		identifier := p.tok
		p.consume()
		return &ast.PseudoClassSelector{Parent: parent, Colon: colon, Identifier: identifier}
	} else {
		return p.consumeCombinatorSelector(parent)
	}
}

func (p *Parser) consumePseudoElementSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming pseudo element selector")

	if p.peekIs(token.KIND_COLON) &&
		p.peekNIs(2, token.KIND_IDENTIFIER) &&
		p.peekNIs(3, token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		colonOne := p.tok
		p.consume()
		colonTwo := p.tok
		p.consume()
		identifier := p.tok
		p.consume()
		return &ast.PseudoElementSelector{
			Parent:     parent,
			ColonOne:   colonOne,
			ColonTwo:   colonTwo,
			Identifier: identifier,
		}
	} else {
		return p.consumeCombinatorSelector(parent)
	}
}

func (p *Parser) consumeCombinatorSelector(parent *ast.StyleRule) ast.SelectorNode {
	// NOTE: DEBUG
	fmt.Println("consuming combinator selector")

	pos := p.pos
	for !p.atEOF() &&
		!p.tokIs(token.KIND_COMMA, token.KIND_CURLY_BRACKET_OPEN) {
		p.consume()
	}
	return &ast.CombinatorSelector{
		Parent: parent,
		Tokens: p.slice(pos, p.pos),
	}
}

func (p *Parser) consumeDeclarations(parent *ast.StyleRule) {
	// NOTE: DEBUG
	fmt.Println("consuming declarations")

	for !p.atEOF() && !p.tokIs(token.KIND_CURLY_BRACKET_CLOSE) {
		parent.Declarations = append(parent.Declarations, p.consumeDeclaration(parent))
	}
}

func (p *Parser) consumeDeclaration(parent *ast.StyleRule) ast.DeclarationNode {
	// NOTE: DEBUG
	fmt.Println("consuming declaration")

	if !p.tokIs(token.KIND_IDENTIFIER) {
		panic(p.unexpectedToken(p.tok, "identifier"))
	}
	if !types.IsProperty(p.tok.Literal()) {
		panic(p.unexpectedToken(p.tok, "property"))
	}

	property := p.tok
	p.consume() // consume identifier

	if !p.tokIs(token.KIND_COLON) {
		panic(p.unexpectedToken(p.tok, ":"))
	}
	colon := p.tok
	p.consume() // consume :

	if !p.tokIs(
		token.KIND_IDENTIFIER,
		token.KIND_NUMBER,
		token.KIND_HASH,
		token.KIND_DIMENSION,
		token.KIND_PERCENTAGE,
	) {
		panic(p.unexpectedToken(p.tok, "value"))
	}
	pos := p.pos
	values := p.consumeDeclarationValueList(pos) // consume any value until ;

	if !p.peekIs(token.KIND_SEMICOLON) {
		panic(p.unexpectedToken(p.tok, ";"))
	}
	p.consume() // consume ;
	semicolon := p.tok
	p.consume() // advance to next token

	return &ast.Declaration{
		Property:  property,
		Colon:     colon,
		Values:    values,
		Semicolon: semicolon,
	}
}

func (p *Parser) consumeDeclarationValueList(start int) []*token.Token {
	for !p.atEOF() && !p.peekIs(token.KIND_SEMICOLON) {
		p.consume()
	}
	return p.slice(start, p.cursor)
}

func (p *Parser) consumeAtRule(parent ast.RuleNode) ast.AtRuleNode {
	// NOTE: DEBUG
	fmt.Println("consume at rule")

	switch p.tok.Literal().String() {
	case "media", "supports":
		return p.consumeMediaOrSupportsAtRule(parent)
	default:
		return nil
	}
}

func (p *Parser) consumeMediaOrSupportsAtRule(parent ast.RuleNode) *ast.MediaOrSupportsAtRule {
	// NOTE: DEBUG
	fmt.Println("consume media at rule")

	a := &ast.MediaOrSupportsAtRule{
		Parent:    parent,
		AtKeyword: p.tok,
	}
	p.consume() // consume the at keyword

	if !p.tokIs(token.KIND_IDENTIFIER, token.KIND_PARENTHESIS_OPEN) {
		panic(p.unexpectedToken(p.tok, "identifier or ("))
	}

	pos := p.pos
	condition := p.consumeMediaOrSupportsAtRuleCondition(pos)
	a.Condition = condition

	p.consume() // consume '{'

	a.StyleRules = append(a.StyleRules, p.consumeStyleRule(nil))

	p.consume() // consume '}' to close the rule

	return a
}

func (p *Parser) consumeMediaOrSupportsAtRuleCondition(start int) []*token.Token {
	// NOTE: DEBUG
	fmt.Println("consume media or supports at rule condition")
	for !p.atEOF() && !p.tokIs(token.KIND_CURLY_BRACKET_OPEN) {
		p.consume()
	}
	return p.slice(start, p.pos)
}

func (p *Parser) atEOF() bool {
	return p.tok.Kind() == token.KIND_EOF
}

func (p *Parser) emit(node ast.Node) {
	p.stylesheet.Nodes = append(p.stylesheet.Nodes, node)
}

func (p *Parser) peek() *token.Token {
	return p.peekN(1)
}

func (p *Parser) tokIs(kind token.Kind, additional ...token.Kind) bool {
	return p.kindIs(p.tok, kind, additional...)
}

func (p *Parser) peekIs(kind token.Kind, additional ...token.Kind) bool {
	return p.kindIs(p.peek(), kind, additional...)
}

func (p *Parser) peekNIs(offset int, kind token.Kind, additional ...token.Kind) bool {
	return p.kindIs(p.peekN(offset), kind, additional...)
}

func (p *Parser) kindIs(tok *token.Token, kind token.Kind, additional ...token.Kind) bool {
	needle := tok.Kind()
	haystack := append(additional, kind)
	for _, piece := range haystack {
		if needle == piece {
			return true
		}
	}
	return false
}

func (p *Parser) peekN(offset int) *token.Token {
	if len(p.tokens) == 0 {
		return token.New(token.KIND_EOF, types.Literal(""), types.NewSourceLocation(0, 0, 0))
	}
	if p.pos+offset >= len(p.tokens) {
		return p.tokens[len(p.tokens)-1]
	} else {
		return p.tokens[p.pos+offset]
	}
}

func (p *Parser) slice(start, end int) []*token.Token {
	return p.tokens[start:end]
}

func (p *Parser) skipWhitespace() {
	for p.tok.Kind() == token.KIND_WHITESPACE {
		p.consume()
	}
}

func (p *Parser) unexpectedToken(got *token.Token, expected string) string {
	gotText := got.Literal().String()
	if gotText == "\x00" {
		return "unexpected end of input"
	}
	if expected == "" {
		return fmt.Sprintf(
			"\n\tunexpected token: %q\n\tposition: line %d, col %d\n",
			got.Literal().String(),
			got.SourceLocation().Row(),
			got.SourceLocation().Start(),
		)
	} else {
		return fmt.Sprintf(
			"\n\tunexpected token: %q\n\texpected: %s\n\tposition: line %d, col %d\n",
			got.Literal().String(),
			expected,
			got.SourceLocation().Row(),
			got.SourceLocation().Start(),
		)
	}
}
