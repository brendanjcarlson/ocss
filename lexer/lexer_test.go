package lexer

import (
	"fmt"
	"testing"

	"github.com/brendanjcarlson/ocss/testutils"
	"github.com/brendanjcarlson/ocss/token"
)

func Test_Lexer(t *testing.T) {
	contents := testutils.MustReadFile(t, "../testdata/simple.ocss")

	l := New(string(contents))

	tokens, elapsed := l.TokenizeTimed()

	for _, tok := range tokens {
		if tok.Kind() != token.KIND_WHITESPACE {
			if tok.Literal().IsKeyword() {
				fmt.Println(tok, "keyword")
			} else if tok.Literal().IsAtKeyword() {
				fmt.Println(tok, "element")
			} else if tok.Literal().IsProperty() {
				fmt.Println(tok, "property")
			} else if tok.Literal().IsFunction() {
				fmt.Println(tok, "function")
			} else if tok.Literal().IsElement() {
				fmt.Println(tok, "element")
			} else if tok.Literal().IsPseudoElement() {
				fmt.Println(tok, "pseudoelement")
			} else if tok.Literal().IsPseudoClass() {
				fmt.Println(tok, "pseudoclass")
			} else {
				fmt.Println(tok)
			}
		}
	}

	fmt.Println(elapsed)
}
