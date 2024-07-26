package types

import (
	"unicode"
)

type Char rune

// String implements fmt.Stringer.
func (c Char) String() string {
	return string(rune(c))
}

func (c Char) Value() Literal {
	return Literal(string(rune(c)))
}

func (c Char) IsIdentifierStart() bool {
	r := rune(c)
	return c.IsIdentifierLike() ||
		r == '.' ||
		r == '#' ||
		r == '['
}

func (c Char) IsIdentifierLike() bool {
	r := rune(c)
	return unicode.IsLetter(r) ||
		unicode.IsDigit(r) ||
		r == '-' ||
		r == '_'
}

func (c Char) IsEOF() bool {
	return rune(c) == '\x00'
}

func (c Char) IsLetter() bool {
	return unicode.IsLetter(rune(c))
}

func (c Char) IsDigit() bool {
	return unicode.IsDigit(rune(c))
}

func (c Char) IsSpace() bool {
	return unicode.IsSpace(rune(c))
}

func (c Char) IsAt() bool {
	return rune(c) == '@'
}

func (c Char) IsPeriod() bool {
	return rune(c) == '.'
}

func (c Char) IsComma() bool {
	return rune(c) == ','
}

func (c Char) IsCurlyBracketOpen() bool {
	return rune(c) == '{'
}

func (c Char) IsCurlyBracketClose() bool {
	return rune(c) == '}'
}

func (c Char) IsSquareBracketOpen() bool {
	return rune(c) == '['
}

func (c Char) IsDollar() bool {
	return rune(c) == '$'
}

func (c Char) IsSquareBracketClose() bool {
	return rune(c) == ']'
}

func (c Char) IsParenthesisOpen() bool {
	return rune(c) == '('
}

func (c Char) IsParenthesisClose() bool {
	return rune(c) == ')'
}

func (c Char) IsColon() bool {
	return rune(c) == ':'
}

func (c Char) IsSemicolon() bool {
	return rune(c) == ';'
}

func (c Char) IsHexValue() bool {
	r := rune(c)
	return unicode.IsDigit(r) ||
		('a' <= r && r <= 'f') ||
		('A' <= r && r <= 'F')
}

func (c Char) IsForwardSlash() bool {
	return rune(c) == '/'
}

func (c Char) IsAsterisk() bool {
	return rune(c) == '*'
}

func (c Char) IsHyphen() bool {
	return rune(c) == '-'
}

func (c Char) IsPlus() bool {
	return rune(c) == '+'
}

func (c Char) IsPercent() bool {
	return rune(c) == '%'
}

func (c Char) IsNewline() bool {
	return rune(c) == '\n'
}

func (c Char) IsAmpersand() bool {
	return rune(c) == '&'
}

func (c Char) IsEqual() bool {
	return rune(c) == '='
}

func (c Char) IsBang() bool {
	return rune(c) == '!'
}

func (c Char) IsHash() bool {
	return rune(c) == '#'
}
