package lexer

import (
	"testing"

	"github.com/jun06t/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectType    token.TokenType
		expectLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectType {
			t.Fatalf("tests[%d] - tokentype wrong.expected=%q, got=%q", i, tt.expectType, tok.Type)
		}
		if tok.Literal != tt.expectLiteral {
			t.Fatalf("tests[%d] - literal wrong.expected=%q, got=%q", i, tt.expectLiteral, tok.Literal)
		}
	}
}
