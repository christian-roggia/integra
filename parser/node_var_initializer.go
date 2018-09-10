package parser

import (
	"fmt"
	"strings"
)

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

func (n *VariableInitializerNode) ToGolang(indent int) string {
	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%s%s := %s", i, n.Name, n.Expression.ToGolang(0))
}

func (n *VariableInitializerNode) ToC(indent int) string {
	i := strings.Repeat(" ", indent*CIndent)
	return fmt.Sprintf("%sint64_t %s = %s;", i, n.Name, n.Expression.ToC(0))
}
