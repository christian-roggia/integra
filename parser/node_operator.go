package parser

import (
	"fmt"
)

type OperatorNode struct {
	Type     NodeType `json:"type"`
	Operator string   `json:"op"`
}

func newOperatorNode(op string) *OperatorNode {
	return &OperatorNode{Type: NodeOperator, Operator: op}
}

func (n *OperatorNode) String() string {
	return n.Operator
}

func (n *OperatorNode) ToGolang(indent int) string {
	return fmt.Sprintf("%s", n.Operator)
}

func (n *OperatorNode) ToC(indent int) string {
	return fmt.Sprintf("%s", n.Operator)
}
