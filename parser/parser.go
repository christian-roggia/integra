package parser

import (
	"fmt"

	"github.com/christian-roggia/integra/lexer"
)

type Parser struct {
	lex *lexer.Lexer

	backup *lexer.Token
}

func NewParser(name, input string) *Parser {
	return &Parser{
		lex: lexer.Lex(name, input),
	}
}

func (parser *Parser) Parse() error {
	return parser.walk()
}

func (parser *Parser) next() lexer.Token {
	if parser.backup != nil {
		next := parser.backup
		parser.backup = nil

		return *next
	}

	return <-parser.lex.Get()
}

func (parser *Parser) peek() lexer.Token {
	if parser.backup != nil {
		return *parser.backup
	}

	next := parser.next()
	parser.backup = &next

	return next
}

func (parser *Parser) walk() error {
	for {
		item := parser.next()
		switch item.Type {
		case lexer.TokenComment:
			// TODO.
		case lexer.TokenKeyword:
			if item.Value != string(lexer.KeywordEquation) {
				return fmt.Errorf("unexpected keyword %s", item.Value)
			}

			if err := parser.walkDeclarator(); err != nil {
				return err
			}

			if err := parser.walkCompoundStatement(); err != nil {
				return err
			}
		case lexer.TokenEOF:
			return nil
		default:
			return fmt.Errorf("unexpected token %s", item.Type.String())
		}
	}
}

func (parser *Parser) walkAssignmentStatement() error {
	context := "clause assignment statement"

	next := parser.next()
	if next.Type != lexer.TokenIdentifier {
		return fmt.Errorf("found '%s', expected identifier in %s", next.Type, context)
	}

	next = parser.next()
	if next.Type != lexer.TokenOperator || (next.Value != ":=" && next.Value != "=") {
		return fmt.Errorf("found '%s', expected assigner in %s", next.Type, context)
	}

	if err := parser.walkExpression(true); err != nil {
		return err
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != ";" {
		return fmt.Errorf("found '%s', expected identifier in %s", next.Type, context)
	}

	return nil
}

func (parser *Parser) walkDeclarator() error {
	context := "clause declarator"

	next := parser.next()
	if next.Type != lexer.TokenIdentifier {
		return fmt.Errorf("found %s, expected identifier in %s", next.String(), context)
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != "(" {
		return fmt.Errorf("found %s, expected parenthesis '(' in %s", next.String(), context)
	}

	peek := parser.peek()
	if peek.Type != lexer.TokenSeparator || peek.Value != ")" {
		if err := parser.walkIdentifierList(); err != nil {
			return err
		}
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != ")" {
		return fmt.Errorf("found %s, expected parenthesis ')' in %s", next.String(), context)
	}

	return nil
}

func (parser *Parser) walkIdentifierList() error {
	context := "clause identifier list"

	next := parser.next()
	if next.Type != lexer.TokenIdentifier {
		return fmt.Errorf("found %s, expected identifier in %s", next.String(), context)
	}

	peek := parser.peek()
	if peek.Type == lexer.TokenSeparator && peek.Value == "," {
		parser.next()
		return parser.walkIdentifierList()
	}

	return nil
}

func (parser *Parser) walkCompoundStatement() error {
	context := "clause compound statement"

	next := parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != "{" {
		return fmt.Errorf("found %s, expected parenthesis '{'in %s", next.String(), context)
	}

	for {
		peek := parser.peek()
		switch peek.Type {
		case lexer.TokenKeyword:
			if peek.Value == string(lexer.KeywordPrint) || peek.Value == string(lexer.KeywordWrite) {
				parser.next()
				if err := parser.walkStringArgumentList(); err != nil {
					return err
				}

				next = parser.next()
				if next.Type != lexer.TokenSeparator || next.Value != ";" {
					return fmt.Errorf("found '%s', expected ';' in %s", next.Value, context)
				}
			}
			if peek.Value == string(lexer.KeywordIf) {
				if err := parser.walkSelectionStatement(); err != nil {
					return err
				}
			}
			if peek.Value == string(lexer.KeywordReturn) {
				if err := parser.walkJumpStatement(); err != nil {
					return err
				}
			}
		case lexer.TokenIdentifier:
			if err := parser.walkAssignmentStatement(); err != nil {
				return err
			}
		default:
			goto FINISH
		}
	}

FINISH:
	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != "}" {
		return fmt.Errorf("found %s, expected parenthesis '}' in %s", next.String(), context)
	}
	return nil
}

func (parser *Parser) walkSelectionStatement() error {
	context := "clause selection statement"

	next := parser.next()
	if next.Type != lexer.TokenKeyword || next.Value != string(lexer.KeywordIf) {
		return fmt.Errorf("found '%s', expected keyword 'if' in %s", next.Value, context)
	}

	if err := parser.walkExpression(true); err != nil {
		return err
	}

	return parser.walkCompoundStatement()
}

func (parser *Parser) walkJumpStatement() error {
	context := "clause jump statement"

	next := parser.next()
	if next.Type != lexer.TokenKeyword || next.Value != string(lexer.KeywordReturn) {
		return fmt.Errorf("found '%s', expected keyword 'return' in %s", next.Value, context)
	}

	if err := parser.walkExpression(false); err != nil {
		return err
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != ";" {
		return fmt.Errorf("found '%s', expected ';' in %s", next.Value, context)
	}
	return nil
}

func isComparisonOperator(s string) bool {
	return s == "==" || s == "<=" || s == ">=" || s == ">" || s == "<"
}

func isArithmeticOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func (parser *Parser) walkExpression(isComparisonAllowed bool) error {
	context := "clause expression"

	if err := parser.walkPrimaryExpression(isComparisonAllowed); err != nil {
		return err
	}

	peek := parser.peek()
	switch peek.Type {
	case lexer.TokenOperator:
		if (isComparisonAllowed && !isComparisonOperator(peek.Value)) &&
			!isArithmeticOperator(peek.Value) {
			return fmt.Errorf("unexpected operator '%s' in %s", peek.Value, context)
		}

		parser.next()
		return parser.walkExpression(isComparisonAllowed)
	default:
		return nil
	}
}

func (parser *Parser) walkPrimaryExpression(isComparisonAllowed bool) error {
	context := "clause primary expression"

	peek := parser.peek()
	switch peek.Type {
	case lexer.TokenIdentifier:
		parser.next()
		peek := parser.peek()
		if peek.Type == lexer.TokenSeparator && peek.Value == "(" {
			return parser.walkArgumentList()
		}
	case lexer.TokenNumber:
		parser.next()
	case lexer.TokenSeparator:
		if peek.Value == "(" {
			parser.next()
			if err := parser.walkExpression(isComparisonAllowed); err != nil {
				return err
			}

			next := parser.next()
			if next.Value != ")" {
				return fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
			}
		}
	}

	return nil
}

func (parser *Parser) walkArgumentList() error {
	context := "clause argument list"

	next := parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != "(" {
		return fmt.Errorf("found '%s', expected parenthesis '(' in %s", next.Value, context)
	}

	if err := parser.walkArgumentExpression(); err != nil {
		return err
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != ")" {
		return fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
	}
	return nil
}

func (parser *Parser) walkStringArgumentList() error {
	context := "clause string argument list"

	next := parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != "(" {
		return fmt.Errorf("found '%s', expected parenthesis '(' in %s", next.Value, context)
	}

	next = parser.next()
	if next.Type != lexer.TokenString {
		return fmt.Errorf("found '%s', expected string in %s", next.Value, context)
	}

	peek := parser.peek()
	if peek.Type == lexer.TokenSeparator && peek.Value == "," {
		parser.next()
		if err := parser.walkArgumentExpression(); err != nil {
			return err
		}
	}

	next = parser.next()
	if next.Type != lexer.TokenSeparator || next.Value != ")" {
		return fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
	}
	return nil
}

func (parser *Parser) walkArgumentExpression() error {
	//context := "clause argument expression"

	if err := parser.walkExpression(false); err != nil {
		return err
	}

	peek := parser.peek()
	if peek.Type == lexer.TokenSeparator && peek.Value == "," {
		parser.next()
		return parser.walkArgumentExpression()
	}

	return nil
}
