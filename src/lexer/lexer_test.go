package lexer

import (
	"monkeylanguage.ayrtonguttier.com.br/token"
	"testing"
)

type tokenValues struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `

    let five = 5;

    let ten = 10;

    let add = fn(x, y){
        x + y;
    };

    let result = add(five, ten);

    `

	tests := []tokenValues{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},

		//	{token.PLUS, "+"},
		//	{token.LPAREN, "("},
		//	{token.RPAREN, ")"},
		//	{token.LBRACE, "{"},
		//	{token.RBRACE, "}"},
		//	{token.COMMA, ","},
		//	{token.SEMICOLON, ";"},
		//	{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] wrong token type. expected: %q, got: %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] wrong literal. expected: %q, got: %q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
