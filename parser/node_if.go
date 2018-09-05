package parser

import (
	"fmt"
	"strings"
)

type IfNode struct {
	Type       NodeType        `json:"type"`
	Expression *ExpressionNode `json:"expr"`
	Block      []Node          `json:"block"`
}

func newIfNode(expr *ExpressionNode) *IfNode {
	return &IfNode{Type: NodeIf, Expression: expr}
}

func (n *IfNode) append(stmts ...Node) {
	n.Block = append(n.Block, stmts...)
}

func (n *IfNode) String() string {
	return ""
}

func (n *IfNode) ToGolang() string {
	var stmts []string
	for _, stmt := range n.Block {
		stmts = append(stmts, stmt.ToGolang())
	}

	return fmt.Sprintf("if %s {\n%s\n}", n.Expression.ToGolang(), strings.Join(stmts, "\n"))
}
