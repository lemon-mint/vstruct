package parser

import (
	"fmt"

	"github.com/lemon-mint/vstruct/ast"
	"github.com/lemon-mint/vstruct/lexer"
	"github.com/lemon-mint/vstruct/parser/utils"
)

type Parser struct {
	filename string

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
		lexer:    lexer,
		filename: lexer.Filename,
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) expect(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) parseStruct() (*ast.Node, error) {
	var node *ast.Node = ast.NewNode(ast.NodeType_STRUCT)
	if !p.expect(lexer.TOKEN_STRUCT) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_STRUCT, p.curToken.Type))
	}
	node.Struct = &ast.Struct{}
	p.nextToken()
	if !p.expect(lexer.TOKEN_IDENTIFIER) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
	}
	node.Name = p.curToken.Literal
	p.nextToken()
	if !p.expect(lexer.TOKEN_OPEN_BRACE) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_OPEN_BRACE, p.curToken.Type))
	}
	p.nextToken()
	for {
		var field *ast.Field = &ast.Field{}
		if !p.expect(lexer.TOKEN_IDENTIFIER) {
			return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
		}
		field.StrType = p.curToken.Literal
		p.nextToken()
		if !p.expect(lexer.TOKEN_IDENTIFIER) {
			return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
		}
		field.Name = p.curToken.Literal
		p.nextToken()
		if !p.expect(lexer.TOKEN_SEMICOLON) {
			return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_SEMICOLON, p.curToken.Type))
		}
		p.nextToken()
		node.Struct.Fields = append(node.Struct.Fields, field)
		if p.expect(lexer.TOKEN_CLOSE_BRACE) {
			p.nextToken()
			break
		}
	}
	return node, nil
}

func (p *Parser) parseAlias() (*ast.Node, error) {
	var node *ast.Node = ast.NewNode(ast.NodeType_ALIAS)
	if !p.expect(lexer.TOKEN_ALIAS) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_ALIAS, p.curToken.Type))
	}
	node.Alias = &ast.Alias{}
	p.nextToken()
	if !p.expect(lexer.TOKEN_IDENTIFIER) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
	}
	node.Name = p.curToken.Literal
	p.nextToken()
	if !p.expect(lexer.TOKEN_EQUAL) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_EQUAL, p.curToken.Type))
	}
	p.nextToken()
	if !p.expect(lexer.TOKEN_IDENTIFIER) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
	}
	node.Alias.StrType = p.curToken.Literal
	p.nextToken()
	if !p.expect(lexer.TOKEN_SEMICOLON) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_SEMICOLON, p.curToken.Type))
	}
	return node, nil
}

func (p *Parser) parseEnum() (*ast.Node, error) {
	var node *ast.Node = ast.NewNode(ast.NodeType_ENUM)
	if !p.expect(lexer.TOKEN_ENUM) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_ENUM, p.curToken.Type))
	}
	node.Enum = &ast.Enum{}
	p.nextToken()
	if !p.expect(lexer.TOKEN_IDENTIFIER) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
	}
	node.Name = p.curToken.Literal
	p.nextToken()
	if !p.expect(lexer.TOKEN_OPEN_BRACE) {
		return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_OPEN_BRACE, p.curToken.Type))
	}
	p.nextToken()
l:
	for {
		if !p.expect(lexer.TOKEN_IDENTIFIER) {
			return nil, fmt.Errorf(utils.Expected(lexer.TOKEN_IDENTIFIER, p.curToken.Type))
		}
		node.Enum.Enums = append(node.Enum.Enums, p.curToken.Literal)
		p.nextToken()
		switch p.curToken.Type {
		case lexer.TOKEN_CLOSE_BRACE:
			p.nextToken()
			break l
		case lexer.TOKEN_COMMA:
			p.nextToken()
			if p.expect(lexer.TOKEN_SEMICOLON) {
				p.nextToken()
			}
		case lexer.TOKEN_SEMICOLON:
			p.nextToken()
		default:
			return nil, fmt.Errorf(utils.Unexpected(p.curToken.Type))
		}
	}
	return node, nil
}

func (p *Parser) Parse() (*ast.File, error) {
	file := ast.NewFile(p.filename)

	for p.curToken.Type != lexer.TOKEN_EOF {
		switch p.curToken.Type {
		case lexer.TOKEN_SEMICOLON:
			p.nextToken()
		case lexer.TOKEN_STRUCT:
			s, err := p.parseStruct()
			if err != nil {
				return nil, err
			}
			file.Nodes = append(file.Nodes, s)
		case lexer.TOKEN_ALIAS:
			a, err := p.parseAlias()
			if err != nil {
				return nil, err
			}
			file.Nodes = append(file.Nodes, a)
		case lexer.TOKEN_ENUM:
			e, err := p.parseEnum()
			if err != nil {
				return nil, err
			}
			file.Nodes = append(file.Nodes, e)
		default:
			return nil, fmt.Errorf(utils.Unexpected(p.curToken.Type))
		}
	}

	return file, nil
}
