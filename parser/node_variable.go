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

func (n *VariableNode) ToGolang(indent int) string {
	return fmt.Sprintf("%s int64", n.Variable)
}

func (n *VariableNode) ToC(indent int) string {
	return fmt.Sprintf("int64_t %s", n.Variable)
}
