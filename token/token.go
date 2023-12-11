package token

type TokenType string

type Token struct {
	Type       TokenType
	Literal    string
	LineNumber int
	FileName   string
}

const (
	STRING              = "STRING"
	ILLEGAL             = "ILLEGAL"
	EOF                 = "EOF"
	ID                  = "ID"
	INT                 = "INT"
	ASSIGN              = "="
	PLUS                = "+"
	COMMA               = ","
	SEMICOLON           = ";"
	LPARENTHESES        = "("
	RPARENTHESES        = ")"
	LBRACE              = "{"
	RBRACE              = "}"
	FUNCTION            = "FUNCTION"
	DEF                 = "DEF"
	RETURN              = "RETURN"
	IF                  = "IF"
	ELSE                = "ELSE"
	GREATER             = ">"
	SMALLER             = "<"
	GREATEREQUAL        = ">="
	SMALLEREQUAL        = "<="
	MINUS               = "-"
	BANG                = "!"
	SLASH               = "/"
	ASTERISK            = "*"
	TRUE                = "TRUE"
	FALSE               = "FALSE"
	ELIF                = "ELIF"
	EQUALITY            = "=="
	NOT_EQUALITY_SIMPLE = "!="
	NOT_EQUALITY_SIGNS  = "<>"
	SHORT_MULTIPLY      = "*="
	SHORT_PLUS          = "+="
	SHORT_DIVISION      = "/="
	SHORT_MINUS         = "-="
	LBRACKET            = "["
	RBRACKET            = "]"
)

var keywords map[string]TokenType = map[string]TokenType{
	"fun":    FUNCTION,
	"def":    DEF,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"elif":   ELIF,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(id string) TokenType {
	if token, ok := keywords[id]; ok {
		return token
	}
	return ID
}
