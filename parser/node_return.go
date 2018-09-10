package parser

import (
	"fmt"
	"strings"
)

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

func (n *ReturnNode) ToGolang(indent int) string {
	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%sreturn %s", i, n.Expression.ToGolang(0))
}

func (n *ReturnNode) ToC(indent int) string {
	i := strings.Repeat(" ", indent*CIndent)
	return fmt.Sprintf("%sreturn %s;", i, n.Expression.ToC(0))
}
