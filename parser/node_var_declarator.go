package parser

import (
	"fmt"
	"strings"
)

type VariableDeclaratorNode struct {
	Type       NodeType        `json:"type"`
	Name       string          `json:"name"`
	Expression *ExpressionNode `json:"expr"`
}

func newVariableDeclaratorNode(name string, expr *ExpressionNode) *VariableDeclaratorNode {
	return &VariableDeclaratorNode{Type: NodeVariableDeclarator, Name: name, Expression: expr}
}

func (n *VariableDeclaratorNode) String() string {
	return ""
}

func (n *VariableDeclaratorNode) ToGolang(indent int) string {
	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%s%s = %s", i, n.Name, n.Expression.ToGolang(0))
}
