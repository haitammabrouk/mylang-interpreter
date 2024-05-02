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
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	NULL = "NULL"

	// operators
	PLUS = "+"
	ASSIGN = "="
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	BANG = "!"

	//Comparison operators
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="
	LE = "<="
	GE = ">="

	//Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"fn" : FUNCTION,
	"let" : LET,
	"if" : IF,
	"else" : ELSE,
	"true" : TRUE,
	"false" : FALSE,
	"return" : RETURN,
	"null" : NULL,
}

// a function to check if the literal is either an identifier or a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}