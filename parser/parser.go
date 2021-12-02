package parser

import "github.com/lemon-mint/vstruct/lexer"

type Parser struct {
	lexer *lexer.Lexer

	curToken  lexer.Token
	peekToken lexer.Token
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func New(lexer *lexer.Lexer) *Parser {
	p := &Parser{
		lexer: lexer,
	}

	p.nextToken()
	p.nextToken()

	return p
}
