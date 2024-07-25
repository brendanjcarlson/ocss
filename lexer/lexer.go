package lexer

import (
	"time"

	"github.com/brendanjcarlson/ocss/token"
)

type Lexer struct {
	input  string
	cursor int
	pos    int
	row    int
	bol    int
	char   token.Char
	tokens []*token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input, 0, 0, 1, 0, 0, make([]*token.Token, 0)}
	l.consume()
	return l
}

func (l *Lexer) Tokenize() []*token.Token {
main:
	for {
		switch {
		case l.char.IsEOF():
			l.emit(token.New(token.KIND_EOF, l.char.Value(), l.position()))
			break main

		case l.char.IsDigit():
			kind, literal, position := l.consumeNumeric(l.pos)
			l.emit(token.New(kind, literal, position))

		case l.char.IsLetter():
			literal, position := l.consumeIdentifierLike(l.pos)
			l.emit(token.New(token.KIND_IDENTIFIER, literal, position))

		case l.char.IsAt():
			l.consume() // consume '@'
			pos := l.pos
			if l.char.IsIdentifierLike() { // if next char is not identifer like, syntax error
				for l.char.IsIdentifierLike() {
					l.consume()
				}
				l.emit(token.New(token.KIND_AT_KEYWORD, l.slice(pos, l.pos), l.positionFrom(pos)))
			} else {
				l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))
			}

		case l.char.IsAmpersand():
			l.emit(token.New(token.KIND_IDENTIFIER, l.char.Value(), l.position()))

		case l.char.IsPeriod():
			switch {
			case l.peek().IsDigit():
				pos := l.pos
				l.consume()
				kind, literal, position := l.consumeNumeric(pos)
				l.emit(token.New(kind, literal, position))

			case l.peek().IsIdentifierLike():
				pos := l.pos
				literal, position := l.consumeIdentifierLike(pos)
				l.emit(token.New(token.KIND_IDENTIFIER, literal, position))

			default:
				l.emit(token.New(token.KIND_ILLEGAL, l.char.Value(), l.position()))
			}

		case l.char.IsHash():
			l.emit(token.New(token.KIND_HASH, l.char.Value(), l.position()))

		case l.char.IsHyphen():
			switch {
			case l.peek().IsDigit(), l.peek().IsPeriod():
				pos := l.pos
				l.consume()
				kind, literal, position := l.consumeNumeric(pos)
				l.emit(token.New(kind, literal, position))

			case l.peek().IsSpace():
				l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))

			case l.peek().IsLetter():
				pos := l.pos
				l.consume()
				literal, position := l.consumeIdentifierLike(pos)
				l.emit(token.New(token.KIND_IDENTIFIER, literal, position))

			case l.peek().IsHyphen():
				literal, position := l.consumeIdentifierLike(l.pos)
				l.emit(token.New(token.KIND_IDENTIFIER, literal, position))

			default:
				l.emit(token.New(token.KIND_ILLEGAL, l.char.Value(), l.position()))
			}

		case l.char.IsPlus():
			l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))

		case l.char.IsCurlyBracketOpen():
			l.emit(token.New(token.KIND_CURLY_BRACKET_OPEN, l.char.Value(), l.position()))

		case l.char.IsCurlyBracketClose():
			l.emit(token.New(token.KIND_CURLY_BRACKET_CLOSE, l.char.Value(), l.position()))

		case l.char.IsParenthesisOpen():
			l.emit(token.New(token.KIND_PARENTHESIS_OPEN, l.char.Value(), l.position()))

		case l.char.IsParenthesisClose():
			l.emit(token.New(token.KIND_PARENTHESIS_CLOSE, l.char.Value(), l.position()))

		case l.char.IsComma():
			l.emit(token.New(token.KIND_COMMA, l.char.Value(), l.position()))

		case l.char.IsSemicolon():
			l.emit(token.New(token.KIND_SEMICOLON, l.char.Value(), l.position()))

		case l.char.IsColon():
			l.emit(token.New(token.KIND_COLON, l.char.Value(), l.position()))

		case l.char.IsEqual():
			switch {
			case l.peek().IsEqual():
				pos := l.pos
				ch := l.char
				l.consume() // consume first '='
				l.emit(token.New(token.KIND_OPERATOR, ch.Value()+l.char.Value(), l.positionFrom(pos)))

			default:
				l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))
			}

		case l.char.IsBang():
			switch {
			case l.peek().IsBang():
				pos := l.pos
				ch := l.char
				l.consume() // consume '!'
				l.emit(token.New(token.KIND_OPERATOR, ch.Value()+l.char.Value(), l.positionFrom(pos)))

			default:
				l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))
			}

		case l.char.IsForwardSlash():
			switch {
			case l.peek().IsAsterisk():
				pos := l.pos
				ch := l.char
				l.consume() // consume '/'
				l.emit(token.New(token.KIND_COMMENT_BLOCK_START, ch.Value()+l.char.Value(), l.positionFrom(pos)))
				l.consume() // consume '*'
				pos = l.pos
			commentblock:
				for { // consume until we see "*/"
					if (l.char.IsAsterisk() && l.peek().IsForwardSlash()) || l.atEOF() {
						break commentblock
					}
					l.consume()
				}
				l.emit(token.New(token.KIND_COMMENT_TEXT, l.slice(pos, l.pos), l.positionFrom(pos)))
				ch = l.char
				pos = l.pos
				l.consume() // consume '/'
				l.emit(token.New(token.KIND_COMMENT_BLOCK_END, ch.Value()+l.char.Value(), l.positionFrom(pos)))

			case l.peek().IsForwardSlash():
				pos := l.pos
				ch := l.char
				l.consume() // consume first '/'
				l.emit(token.New(token.KIND_COMMENT_LINE_START, ch.Value()+l.char.Value(), l.positionFrom(pos)))
				l.consume() // consume second '/'
				pos = l.pos
			commentline:
				for !l.peek().IsNewline() {
					if l.atEOF() {
						break commentline
					}
					l.consume()
				}
				l.emit(token.New(token.KIND_COMMENT_TEXT, l.slice(pos, l.cursor), l.positionFrom(pos)))

			default:
				l.emit(token.New(token.KIND_OPERATOR, l.char.Value(), l.position()))
			}

		case l.char.IsAsterisk():
			pos := l.pos
			ch := l.char
			if l.peek().IsForwardSlash() {
				l.consume()
				l.emit(token.New(token.KIND_COMMENT_BLOCK_END, ch.Value()+l.char.Value(), l.positionFrom(pos)))
			} else {
				l.emit(token.New(token.KIND_OPERATOR, ch.Value(), l.position()))
			}

		case l.char.IsIdentifierLike():
			pos := l.pos
			literal, position := l.consumeIdentifierLike(pos)
			l.emit(token.New(token.KIND_IDENTIFIER, literal, position))

		case l.char.IsSpace():
			// noop

		default:
			l.emit(token.New(token.KIND_ILLEGAL, l.char.Value(), l.position()))
		}

		l.consume()
	}

	return l.tokens
}

func (l *Lexer) TokenizeTimed() ([]*token.Token, string) {
	now := time.Now()
	tokens := l.Tokenize()
	elapsed := time.Since(now)
	return tokens, "\nTOKENIZED IN " + elapsed.String() + "\n"
}

func (l *Lexer) atEOF() bool {
	return l.cursor >= len(l.input)
}

func (l *Lexer) position() token.Position {
	return token.NewPosition(l.row, l.pos-l.bol+1, l.pos-l.bol+1)
}

func (l *Lexer) positionFrom(start int) token.Position {
	return token.NewPosition(l.row, start-l.bol+1, l.pos-l.bol+1)
}

func (l *Lexer) slice(start, end int) token.Literal {
	if end >= len(l.input) {
		end = len(l.input) - 1
	}
	return token.Literal(l.input[start:end])
}

func (l *Lexer) peek() token.Char {
	return l.peekN(0)
}

func (l *Lexer) peekN(offset int) token.Char {
	if l.atEOF() || l.cursor+offset >= len(l.input) {
		return 0
	} else {
		return token.Char(l.input[l.cursor+offset])
	}
}

func (l *Lexer) peekString(length int) string {
	if l.atEOF() || l.cursor+length >= len(l.input) {
		return ""
	} else {
		return l.input[l.cursor : l.cursor+length]
	}
}

func (l *Lexer) consume() {
	if l.atEOF() {
		l.char = 0
	} else {
		l.char = token.Char(l.input[l.cursor])
		if l.char == '\n' {
			l.row += 1
			l.bol = l.cursor + 1
		}
	}
	l.pos = l.cursor
	l.cursor += 1
}

func (l *Lexer) emit(t *token.Token) {
	l.tokens = append(l.tokens, t)
}

func (l *Lexer) consumeNumeric(start int) (token.Kind, token.Literal, token.Position) {
	seenPeriod := false
numeric:
	for {
		if (l.atEOF() || l.peek().IsSemicolon() || l.peek().IsComma()) &&
			(!l.peek().IsDigit() || !l.peek().IsLetter() || !l.peek().IsPeriod() || !l.peek().IsPercent()) {
			break numeric
		}

		if l.peek().IsPeriod() {
			if !seenPeriod {
				seenPeriod = true
			} else {
				l.consume()
				return token.KIND_ILLEGAL, l.slice(start, l.cursor), l.positionFrom(start)
			}
		} else if l.peek().IsPercent() {
			l.consume()
			return token.KIND_PERCENTAGE, l.slice(start, l.cursor), l.positionFrom(start)
		} else if l.peek().IsLetter() {
			for l.peek().IsLetter() {
				l.consume()
			}
			return token.KIND_DIMENSION, l.slice(start, l.cursor), l.positionFrom(start)
		} else {
			l.consume()
		}
	}
	return token.KIND_NUMBER, l.slice(start, l.cursor), l.positionFrom(start)
}

func (l *Lexer) consumeIdentifierLike(start int) (token.Literal, token.Position) {
	for l.peek().IsIdentifierLike() {
		l.consume()
	}
	return l.slice(start, l.cursor), l.positionFrom(start)
}
