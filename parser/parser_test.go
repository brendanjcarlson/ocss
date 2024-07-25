package parser

import (
	"fmt"
	"testing"

	"github.com/brendanjcarlson/ocss/lexer"
	"github.com/brendanjcarlson/ocss/testutils"
)

func Test_Parser(t *testing.T) {
	contents := testutils.MustReadFile(t, "../testdata/simple.ocss")

	l := lexer.New(string(contents))

	tokens := l.Tokenize()
	// for _, tok := range tokens {
	// 	if tok.Kind() != token.KIND_WHITESPACE {
	// 		fmt.Println(tok)
	// 	}
	// }

	p := New(tokens)

	stylesheet := p.ParseStyleSheet()

	fmt.Println(stylesheet.Literal())
}
