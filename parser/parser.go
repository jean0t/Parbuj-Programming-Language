package parser

import (
	"tarbuj/ast"
	"tarbuj/lexer"
	"tarbuj/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

// more
