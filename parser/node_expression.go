package parser

import (
	"fmt"
	"strings"
)

type ExpressionNode struct {
	Type  NodeType `json:"type"`
	Nodes []Node   `json:"children"`
}

func newExpressionNode(children []Node) *ExpressionNode {
	return &ExpressionNode{Type: NodeExpression, Nodes: children}
}

func (expr *ExpressionNode) append(n ...Node) {
	expr.Nodes = append(expr.Nodes, n...)
}

func (expr *ExpressionNode) String() string {
	return ""
}

func (expr *ExpressionNode) ToGolang(indent int) string {
	var s []string
	for _, e := range expr.Nodes {
		s = append(s, e.ToGolang(0))
	}

	return fmt.Sprintf("%s", strings.Join(s, " "))
}
