package lexer

import "github.com/leewei05/monkey/token"

// PROD: consider using io.Reader to handle filename and line numbers
type Lexer struct {
	input string
	// position is the current position in input(points to current char)
	position int
	// readPosition is the current reading position(after current char)
	readPosition int
	// ch is the current char under examination
	ch byte
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func (l *Lexer) NextToken() *token.Token {
	return &token.Token{}
}
