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
		"ERROR",
		"IDENTIFIER",
		"KEYWORD",
		"SEPARATOR",
		"OPERATOR",
		"NUMBER",
		"STRING",
		"COMMENT",
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

	return fmt.Sprintf("%s '%s'", t.Type, t.Value)
}

func (t *Token) IsKeyword(kwType KeywordType) bool {
	if t.Type != TokenKeyword {
		return false
	}

	return t.Value == string(kwType)
}

func (t *Token) IsOperator(opType OperatorType) bool {
	if t.Type != TokenOperator {
		return false
	}

	return t.Value == string(opType)
}

func (t *Token) IsSeparator(sepType SeparatorType) bool {
	if t.Type != TokenSeparator {
		return false
	}

	return t.Value == string(sepType)
}

func (t *Token) IsIdentifier() bool {
	return t.Type == TokenIdentifier
}
