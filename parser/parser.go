package parser

import (
	"fmt"

	"github.com/christian-roggia/integra/lexer"
)

type Parser struct {
	lex    *lexer.Lexer
	backup *lexer.Token

	tr *GlobalNode
}

func NewParser(name, input string) *Parser {
	return &Parser{
		lex: lexer.Lex(name, input),
	}
}

func (parser *Parser) Parse() error {
	return parser.walk()
}

func (parser *Parser) Tree() *GlobalNode {
	return parser.tr
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
	parser.tr = newGlobalNode()

	for {
		next := parser.next()
		switch next.Type {
		case lexer.TokenComment:
			parser.tr.append(newCommentNode(next.Value))
		case lexer.TokenKeyword:
			if !next.IsKeyword(lexer.KwEquation) {
				return fmt.Errorf("unexpected keyword '%s'", next.Value)
			}

			id, args, err := parser.walkDeclarator()
			if err != nil {
				return err
			}

			stmts, err := parser.walkCompoundStatement()
			if err != nil {
				return err
			}

			equationNode := newEquationNode(id, args)
			equationNode.append(stmts...)

			parser.tr.append(equationNode)
		case lexer.TokenEOF:
			return nil
		default:
			return fmt.Errorf("unexpected token %s", next.String())
		}
	}
}

func (parser *Parser) walkAssignmentStatement() (Node, error) {
	context := "clause assignment statement"

	next := parser.next()
	if !next.IsIdentifier() {
		return nil, fmt.Errorf("found '%s', expected identifier in %s", next.Type, context)
	}
	id := next.Value

	next = parser.next()
	if !next.IsOperator(lexer.OpAssignment) && !next.IsOperator(lexer.OpInitialization) {
		return nil, fmt.Errorf("found '%s', expected assigner in %s", next.Type, context)
	}
	kind := next.Value

	expr, err := parser.walkExpression(true)
	if err != nil {
		return nil, err
	}

	next = parser.next()
	if !next.IsSeparator(lexer.SepSemicolon) {
		return nil, fmt.Errorf("found '%s', expected ';' in %s", next.Type, context)
	}

	if kind == lexer.OpInitialization {
		return newVariableInitializerNode(id, newExpressionNode(expr)), nil
	}

	return newVariableDeclaratorNode(id, newExpressionNode(expr)), nil
}

func (parser *Parser) walkDeclarator() (string, []*VariableNode, error) {
	context := "clause declarator"

	next := parser.next()
	if !next.IsIdentifier() {
		return "<invalid>", nil, fmt.Errorf("found '%s', expected identifier in %s", next.Value, context)
	}
	id := next.Value

	next = parser.next()
	if !next.IsSeparator(lexer.SepLeftRoundBracket) {
		return id, nil, fmt.Errorf("found '%s', expected parenthesis '(' in %s", next.Value, context)
	}

	var args []*VariableNode
	peek := parser.peek()
	if !peek.IsSeparator(lexer.SepRightRoundBracket) {
		var err error
		args, err = parser.walkIdentifierList()
		if err != nil {
			return id, nil, err
		}
	}

	next = parser.next()
	if !next.IsSeparator(lexer.SepRightRoundBracket) {
		return id, nil, fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
	}

	return id, args, nil
}

func (parser *Parser) walkIdentifierList() ([]*VariableNode, error) {
	context := "clause identifier list"

	next := parser.next()
	if !next.IsIdentifier() {
		return nil, fmt.Errorf("found '%s', expected identifier in %s", next.Value, context)
	}
	args := []*VariableNode{newVariableNode(next.Value)}

	peek := parser.peek()
	if peek.IsSeparator(lexer.SepComma) {
		parser.next()

		list, err := parser.walkIdentifierList()
		if err != nil {
			return nil, err
		}

		return append(args, list...), nil
	}

	return args, nil
}

func (parser *Parser) walkCompoundStatement() ([]Node, error) {
	context := "clause compound statement"

	next := parser.next()
	if !next.IsSeparator(lexer.SepLeftCurlyBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis '{' in %s", next.Value, context)
	}

	var stmts []Node
	for {
		peek := parser.peek()
		switch peek.Type {
		case lexer.TokenKeyword:
			if peek.IsKeyword(lexer.KwPrint) || peek.IsKeyword(lexer.KwWrite) {
				next := parser.next()
				id := next.Value

				args, err := parser.walkStringArgumentList()
				if err != nil {
					return nil, err
				}

				next = parser.next()
				if !next.IsSeparator(lexer.SepSemicolon) {
					return nil, fmt.Errorf("found '%s', expected ';' in %s", next.Value, context)
				}

				stmts = append(stmts, newCallNode(id, args))
			}
			if peek.Value == string(lexer.KwIf) {
				stmt, err := parser.walkSelectionStatement()
				if err != nil {
					return nil, err
				}
				stmts = append(stmts, stmt)
			}
			if peek.Value == string(lexer.KwReturn) {
				stmt, err := parser.walkJumpStatement()
				if err != nil {
					return nil, err
				}
				stmts = append(stmts, stmt)
			}
		case lexer.TokenIdentifier:
			stmt, err := parser.walkAssignmentStatement()
			if err != nil {
				return nil, err
			}
			stmts = append(stmts, stmt)
		default:
			goto FINISH
		}
	}

FINISH:
	next = parser.next()
	if !next.IsSeparator(lexer.SepRightCurlyBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis '}' in %s", next.Value, context)
	}
	return stmts, nil
}

func (parser *Parser) walkSelectionStatement() (*IfNode, error) {
	context := "clause selection statement"

	next := parser.next()
	if !next.IsKeyword(lexer.KwIf) {
		return nil, fmt.Errorf("found '%s', expected keyword 'if' in %s", next.Value, context)
	}

	expr, err := parser.walkExpression(true)
	if err != nil {
		return nil, err
	}
	stmt := newIfNode(newExpressionNode(expr))

	stmts, err := parser.walkCompoundStatement()
	stmt.append(stmts...)

	return stmt, err
}

func (parser *Parser) walkJumpStatement() (*ReturnNode, error) {
	context := "clause jump statement"

	next := parser.next()
	if !next.IsKeyword(lexer.KwReturn) {
		return nil, fmt.Errorf("found '%s', expected keyword 'return' in %s", next.Value, context)
	}

	expr, err := parser.walkExpression(false)
	if err != nil {
		return nil, err
	}

	next = parser.next()
	if !next.IsSeparator(lexer.SepSemicolon) {
		return nil, fmt.Errorf("found '%s', expected ';' in %s", next.Value, context)
	}
	return newReturnNode(newExpressionNode(expr)), nil
}

func isComparisonOperator(s string) bool {
	return s == "==" || s == "<=" || s == ">=" || s == ">" || s == "<"
}

func isArithmeticOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func (parser *Parser) walkExpression(isComparisonAllowed bool) ([]Node, error) {
	context := "clause expression"

	expr, err := parser.walkPrimaryExpression(isComparisonAllowed)
	if err != nil {
		return nil, err
	}
	if expr == nil {
		return nil, nil
	}
	exprs := []Node{expr}

	peek := parser.peek()
	switch peek.Type {
	case lexer.TokenOperator:
		if (isComparisonAllowed && !isComparisonOperator(peek.Value)) &&
			!isArithmeticOperator(peek.Value) {
			return nil, fmt.Errorf("unexpected operator '%s' in %s", peek.Value, context)
		}

		next := parser.next()
		exprs = append(exprs, newOperatorNode(next.Value))

		expr, err := parser.walkExpression(isComparisonAllowed)
		if err != nil {
			return nil, err
		}
		return append(exprs, expr...), nil
	default:
		return exprs, nil
	}
}

func (parser *Parser) walkPrimaryExpression(isComparisonAllowed bool) (Node, error) {
	context := "clause primary expression"

	peek := parser.peek()
	switch peek.Type {
	case lexer.TokenIdentifier:
		next := parser.next()
		id := next.Value

		peek := parser.peek()
		if peek.IsSeparator(lexer.SepLeftRoundBracket) {
			args, err := parser.walkArgumentList()
			if err != nil {
				return nil, err
			}

			return newCallNode(id, args), nil
		}

		return newNumberNode(id), nil
	case lexer.TokenNumber:
		next := parser.next()
		return newNumberNode(next.Value), nil
	case lexer.TokenSeparator:
		if peek.IsSeparator(lexer.SepLeftRoundBracket) {
			parser.next()

			expr, err := parser.walkExpression(isComparisonAllowed)
			if err != nil {
				return nil, err
			}

			next := parser.next()
			if !next.IsSeparator(lexer.SepRightRoundBracket) {
				return nil, fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
			}
			if len(expr) == 0 {
				return nil, fmt.Errorf("unexpected empty expression '()' in %s", context)
			}

			return newExpressionNode(expr), nil
		}
	}

	return nil, nil
}

func (parser *Parser) walkArgumentList() ([]Node, error) {
	context := "clause argument list"

	next := parser.next()
	if !next.IsSeparator(lexer.SepLeftRoundBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis '(' in %s", next.Value, context)
	}

	args, err := parser.walkArgumentExpression()
	if err != nil {
		return nil, err
	}

	next = parser.next()
	if !next.IsSeparator(lexer.SepRightRoundBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
	}
	return args, nil
}

func (parser *Parser) walkStringArgumentList() ([]Node, error) {
	context := "clause string argument list"

	next := parser.next()
	if !next.IsSeparator(lexer.SepLeftRoundBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis '(' in %s", next.Value, context)
	}

	next = parser.next()
	if next.Type != lexer.TokenString {
		return nil, fmt.Errorf("found '%s', expected string in %s", next.Value, context)
	}
	args := []Node{newStringNode(next.Value)}

	peek := parser.peek()
	if peek.IsSeparator(lexer.SepComma) {
		parser.next()

		arg, err := parser.walkArgumentExpression()
		if err != nil {
			return nil, err
		}
		args = append(args, arg...)
	}

	next = parser.next()
	if !next.IsSeparator(lexer.SepRightRoundBracket) {
		return nil, fmt.Errorf("found '%s', expected parenthesis ')' in %s", next.Value, context)
	}
	return args, nil
}

func (parser *Parser) walkArgumentExpression() ([]Node, error) {
	//context := "clause argument expression"

	expr, err := parser.walkExpression(false)
	if err != nil {
		return nil, err
	}
	if len(expr) == 0 {
		return nil, nil
	}

	exprs := []Node{newExpressionNode(expr)}
	peek := parser.peek()
	if peek.IsSeparator(lexer.SepComma) {
		parser.next()

		expr, err := parser.walkArgumentExpression()
		if err != nil {
			return nil, err
		}
		exprs = append(exprs, expr...)
	}

	return exprs, nil
}
