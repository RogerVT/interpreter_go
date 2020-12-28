package lexer

import (
	"interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASIGN, "="},
		{token.MAS, "+"},
		{token.IPAREN, "("},
		{token.RPAREN, ")"},
		{token.ILLAVE, "{"},
		{token.RLLAVE, "}"},
		{token.COMA, ","},
		{token.SEMICOLON, ";"},
		{token.FIN, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType erroneo, esperaba=%q, recibi=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal erronea, esperaba=%q, recibi=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
