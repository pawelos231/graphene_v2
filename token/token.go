package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" //illegal characters
	EOF     = "EOF"     //end of file

	IDENT = "IDENT"
	INT   = "INT"

	//operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	BANG    = "!"
	SLASH   = "/"
	ASTERIK = "*"
	LT      = "<"
	GT      = ">"
	EQ      = "=="
	NOT_EQ  = "!="
	POWER   = "^"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	//loop constructs
	FOR   = "FOR"
	WHILE = "WHILE"
)

type Token struct {
	Type    TokenType // the type of the token
	Literal string    // the literal value of the token
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"for":    FOR,
	"while":  WHILE,
}

func lookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
