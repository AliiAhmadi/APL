package token

type TokenType string

type Token struct {
	Type       TokenType
	Literal    string
	LineNumber int
	FileName   string
}

const (
	ILLEGAL      = "ILLEGAL"
	EOF          = "EOF"
	ID           = "ID"
	INT          = "INT"
	ASSIGN       = "="
	PLUS         = "+"
	COMMA        = ","
	SEMICOLON    = ";"
	LPARENTHESES = "("
	RPARENTHESES = ")"
	LBRACE       = "{"
	RBRACE       = "}"
	FUNCTION     = "FUNCTION"
	DEF          = "DEF"
	RETURN       = "RETURN"
	IF           = "IF"
	ELSE         = "ELSE"
	GREATER      = ">"
	SMALLER      = "<"
	GREATEREQUAL = ">="
	SMALLEREQUAL = "<="
	MINUS        = "-"
	BANG         = "!"
	SLASH        = "/"
	ASTERISK     = "*"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fun":    FUNCTION,
	"def":    DEF,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdentifier(id string) TokenType {
	if token, ok := keywords[id]; ok {
		return token
	}
	return ID
}
