package lexer

import (
	"fmt"
)

// TokenType enumerates all the valid tokens types.
type TokenType int

// https://en.wikipedia.org/wiki/Lexical_analysis#Token
const (
	TokenEOF TokenType = iota
	TokenError

	TokenIdentifier
	TokenKeyword
	TokenSeparator
	TokenOperator
	TokenNumber
	TokenString
	TokenComment
)

func (tt TokenType) String() string {
	s := []string{
		"EOF",
		"Error",
		"Identifier",
		"Keyword",
		"Separator",
		"Operator",
		"Number",
		"String",
		"Comment",
	}

	if int(tt) < len(s) {
		return s[tt]
	}

	return ""
}

// Token represents a single parsed token.
type Token struct {
	Type  TokenType
	Pos   int
	Value string
}

func (t *Token) String() string {
	if t.Type == TokenEOF {
		return "EOF"
	}

	return fmt.Sprintf("%-12v: %s", t.Type, t.Value)
}
