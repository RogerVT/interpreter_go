package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILEGAL = "ILEGAL"
	FIN    = "FIN"

	//IDENTIFICADORES Y LITERALES
	IDENT = "IDENT"
	INT   = "INT"

	//OPERADORES
	ASIGN = "="
	MAS   = "+"
	MENOS = "-"

	//DELIMITADORES
	COMA      = ","
	SEMICOLON = ";"
	IPAREN    = "("
	RPAREN    = ")"
	ILLAVE    = "{"
	RLLAVE    = "}"

	//DEFINICIONES
	FUNCION = "FUNCION"
	LET     = "LET"
)
