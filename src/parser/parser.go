package parser

import (
	"monkeylanguage.ayrtonguttier.com.br/ast"
	"monkeylanguage.ayrtonguttier.com.br/lexer"
	"monkeylanguage.ayrtonguttier.com.br/token"
)



type Parser struct { 
    l *lexer.Lexer
    curToken token.Token
    peekToken token.Token
}


func New(l *lexer.Lexer) *Parser{
    p := &Parser{l:l}
    p.nextToken()
    p.nextToken()
    return p
}


func (p *Parser) nextToken(){
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}


func (p *Parser) ParseProgram() *ast.Program{
    return nil
}
