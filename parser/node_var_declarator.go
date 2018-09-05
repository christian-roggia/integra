package parser

import "fmt"

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

func (n *VariableDeclaratorNode) ToGolang() string {
	return fmt.Sprintf("%s = %s", n.Name, n.Expression.ToGolang())
}
