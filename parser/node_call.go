package parser

import (
	"fmt"
	"strings"
)

type CallNode struct {
	Type      NodeType `json:"type"`
	Name      string   `json:"name"`
	Arguments []Node   `json:"args"`
}

func newCallNode(name string, args []Node) *CallNode {
	return &CallNode{Type: NodeCall, Name: name, Arguments: args}
}

func (c *CallNode) String() string {
	return fmt.Sprintf("")
}

func (c *CallNode) ToGolang(indent int) string {
	var args []string
	for _, arg := range c.Arguments {
		args = append(args, arg.ToGolang(0))
	}

	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%s%s(%s)", i, c.Name, strings.Join(args, ", "))
}

func (c *CallNode) ToC(indent int) string {
	var args []string
	for _, arg := range c.Arguments {
		args = append(args, arg.ToC(0))
	}

	i := strings.Repeat(" ", indent*CIndent)
	s := fmt.Sprintf("%s%s(%s)", i, c.Name, strings.Join(args, ", "))
	if c.Name == "print" || c.Name == "write" {
		s = fmt.Sprintf("%s;", s)
	}

	return s
}
