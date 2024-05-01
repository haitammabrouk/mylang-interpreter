package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// delimiters
	PLUS = "+"
	ASSIGN = "="

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"

)