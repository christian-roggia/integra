package parser

import "fmt"

type ReturnNode struct {
	Type       NodeType        `json:"type"`
	Expression *ExpressionNode `json:"expr"`
}

func newReturnNode(expr *ExpressionNode) *ReturnNode {
	return &ReturnNode{Type: NodeReturn, Expression: expr}
}

func (n *ReturnNode) String() string {
	return ""
}

func (n *ReturnNode) ToGolang() string {
	return fmt.Sprintf("return %s", n.Expression.ToGolang())
}
