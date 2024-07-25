package token

type Kind int

const (
	KIND_TODO Kind = iota
	KIND_ILLEGAL
	KIND_EOF

	// W3C Tokens
	KIND_IDENTIFIER // a-z A-Z - _
	KIND_AT_KEYWORD // @ followed by a-z A-Z - _ (does not include @)
	KIND_HASH
	KIND_FUNCTION
	KIND_STRING
	KIND_BAD_STRING
	KIND_URL
	KIND_BAD_URL
	KIND_DELIMITER
	KIND_NUMBER
	KIND_PERCENTAGE
	KIND_DIMENSION
	KIND_WHITESPACE
	KIND_COMMENT_TEXT
	KIND_COMMENT_LINE_START
	KIND_COMMENT_BLOCK_START
	KIND_COMMENT_BLOCK_END
	KIND_COLON
	KIND_SEMICOLON
	KIND_COMMA
	KIND_SQUARE_BRACKET_OPEN
	KIND_SQUARE_BRACKET_CLOSE
	KIND_PARENTHESIS_OPEN
	KIND_PARENTHESIS_CLOSE
	KIND_CURLY_BRACKET_OPEN
	KIND_CURLY_BRACKET_CLOSE
	KIND_OPERATOR
)

// String implements fmt.Stringer.
func (k Kind) String() string {
	switch k {
	case KIND_TODO:
		return "TODO"
	case KIND_ILLEGAL:
		return "ILLEGAL"
	case KIND_EOF:
		return "EOF"
	case KIND_IDENTIFIER:
		return "IDENTIFIER"
	case KIND_AT_KEYWORD:
		return "AT_KEYWORD"
	case KIND_HASH:
		return "HASH"
	case KIND_FUNCTION:
		return "FUNCTION"
	case KIND_STRING:
		return "STRING"
	case KIND_BAD_STRING:
		return "BAD_STRING"
	case KIND_URL:
		return "URL"
	case KIND_BAD_URL:
		return "BAD_URL"
	case KIND_DELIMITER:
		return "DELIMITER"
	case KIND_NUMBER:
		return "NUMBER"
	case KIND_PERCENTAGE:
		return "PERCENTAGE"
	case KIND_DIMENSION:
		return "DIMENSION"
	case KIND_WHITESPACE:
		return "WHITESPACE"
	case KIND_COMMENT_TEXT:
		return "COMMENT_TEXT"
	case KIND_COMMENT_LINE_START:
		return "COMMENT_LINE_START"
	case KIND_COMMENT_BLOCK_START:
		return "COMMENT_BLOCK_START"
	case KIND_COMMENT_BLOCK_END:
		return "COMMENT_BLOCK_END"
	case KIND_COLON:
		return "COLON"
	case KIND_SEMICOLON:
		return "SEMICOLON"
	case KIND_COMMA:
		return "COMMA"
	case KIND_SQUARE_BRACKET_OPEN:
		return "SQUARE_BRACKET_OPEN"
	case KIND_SQUARE_BRACKET_CLOSE:
		return "SQUARE_BRACKET_CLOSE"
	case KIND_PARENTHESIS_OPEN:
		return "PARENTHESIS_OPEN"
	case KIND_PARENTHESIS_CLOSE:
		return "PARENTHESIS_CLOSE"
	case KIND_CURLY_BRACKET_OPEN:
		return "CURLY_BRACKET_OPEN"
	case KIND_CURLY_BRACKET_CLOSE:
		return "CURLY_BRACKET_CLOSE"
	case KIND_OPERATOR:
		return "OPERATOR"
	default:
		return "UNKNOWN"
	}
}
