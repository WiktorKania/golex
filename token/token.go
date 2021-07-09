package token

type TokenType int

const (
	FUNCTION TokenType = iota + 130
	RETURN
	IF
	ELSE
	BREAK
	CONTINUE
	FOR
	WHILE
	ALLOC
	DEL
	TYPE
	REF
	CONST
	STRING_LITERAL
	INTEGER
	FLOAT
	HEX
	OCTAL
	BINARY
	ARRAY
	CLASS
	PLUS_ASSIGN
	MINUS_ASSIGN
	MUL_ASSIGN
	DIV_ASSIGN
	MOD_ASSIGN
	LSHIFT_ASSIGN
	RSHIFT_ASSIGN
	AND_ASSIGN
	OR_ASSIGN
	XOR_ASSIGN
	LSHIFT
	RSHIFT
	OP_AND
	OP_OR
	OP_EQ
	OP_NOT_EQ
	OP_GT_EQ
	OP_LT_EQ
	IDENTIFIER
	TRUE
	FALSE
	UNKNOWN
	NONE
	GENERAL_EXPRESSION
	GENERAL_STATEMENT
	OP_PLUS  = '+'
	OP_MINUS = '-'
	OP_DIV   = '/'
)

type Token struct {
	value  string
	line   int
	source string
	_type  TokenType
}

func GetName(_type TokenType) string {
	return "getName(): TODO!"
}

func New(_type TokenType, val string, line int, source string) *Token {
	return &Token{val, line, source, _type}
}
