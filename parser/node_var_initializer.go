package parser

import "fmt"

type VariableInitializerNode struct {
	Type       NodeType        `json:"type"`
	Name       string          `json:"name"`
	Expression *ExpressionNode `json:"expr"`
}

func newVariableInitializerNode(name string, expr *ExpressionNode) *VariableInitializerNode {
	return &VariableInitializerNode{Type: NodeVariableInitializer, Name: name, Expression: expr}
}

func (n *VariableInitializerNode) String() string {
	return ""
}

func (n *VariableInitializerNode) ToGolang() string {
	return fmt.Sprintf("%s := %s", n.Name, n.Expression.ToGolang())
}
