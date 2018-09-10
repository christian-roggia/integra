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

func (n *IfNode) ToGolang(indent int) string {
	var stmts []string
	for _, stmt := range n.Block {
		stmts = append(stmts, stmt.ToGolang(indent+1))
	}

	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%sif %s {\n%s\n%s}", i, n.Expression.ToGolang(0), strings.Join(stmts, "\n"), i)
}

func (n *IfNode) ToC(indent int) string {
	var stmts []string
	for _, stmt := range n.Block {
		stmts = append(stmts, stmt.ToC(indent+1))
	}

	i := strings.Repeat(" ", indent*CIndent)
	return fmt.Sprintf("%sif (%s) {\n%s\n%s}", i, n.Expression.ToC(0), strings.Join(stmts, "\n"), i)
}
