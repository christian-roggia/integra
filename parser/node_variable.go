package parser

import "fmt"

type VariableNode struct {
	Type     NodeType `json:"type"`
	Variable string   `json:"var"`
}

func newVariableNode(v string) *VariableNode {
	return &VariableNode{Type: NodeVariable, Variable: v}
}

func (n *VariableNode) String() string {
	return n.Variable
}

func (n *VariableNode) ToGolang() string {
	return fmt.Sprintf("%s int64", n.Variable)
}
