package parser

import (
	"fmt"
	"strings"
)

type EquationNode struct {
	Type      NodeType        `json:"type"`
	Arguments []*VariableNode `json:"args"`
	Block     []Node          `json:"block"`
	Name      string          `json:"name"`
}

func newEquationNode(name string, args []*VariableNode) *EquationNode {
	return &EquationNode{Type: NodeEquation, Name: name, Arguments: args}
}

func (eq *EquationNode) append(n ...Node) {
	eq.Block = append(eq.Block, n...)
}

func (eq *EquationNode) String() string {
	return ""
}

func (eq *EquationNode) ToGolang() string {
	var args []string
	for _, arg := range eq.Arguments {
		args = append(args, arg.ToGolang())
	}

	var stmts []string
	for _, stmt := range eq.Block {
		stmts = append(stmts, stmt.ToGolang())
	}

	return fmt.Sprintf("func %s(%s) int64 {\n%s\n}", eq.Name,
		strings.Join(args, ", "), strings.Join(stmts, "\n"))
}
