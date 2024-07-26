package lexer

import (
	"fmt"
	"testing"

	"github.com/brendanjcarlson/ocss/testutils"
	"github.com/brendanjcarlson/ocss/token"
	"github.com/brendanjcarlson/ocss/types"
)

func Test_Lexer(t *testing.T) {
	contents := testutils.MustReadFile(t, "../testdata/simple.ocss")

	l := New(string(contents))

	tokens, elapsed := l.TokenizeTimed()

	for _, tok := range tokens {
		if tok.Kind() != token.KIND_WHITESPACE {
			if types.IsAtKeyword(tok.Literal()) {
				fmt.Println(tok, "element")
			} else if types.IsProperty(tok.Literal()) {
				fmt.Println(tok, "property")
			} else if types.IsFunction(tok.Literal()) {
				fmt.Println(tok, "function")
			} else if types.IsHtmlElement(tok.Literal()) {
				fmt.Println(tok, "html element")
			} else if types.IsPseudoElement(tok.Literal()) {
				fmt.Println(tok, "pseudo element")
			} else if types.IsPseudoClass(tok.Literal()) {
				fmt.Println(tok, "pseudo class")
			} else {
				fmt.Println(tok)
			}
		}
	}

	fmt.Println(elapsed)
}
