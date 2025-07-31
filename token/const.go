package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers and literals
	IDENT  = "IDENT"
	INT    = "INT"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ    = "=="
	NOTEQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
