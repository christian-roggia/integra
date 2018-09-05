package parser

import (
	"bytes"
	"fmt"
	"strings"
)

type GlobalNode struct {
	Type  NodeType `json:"type"`
	Nodes []Node   `json:"children"`
}

func newGlobalNode() *GlobalNode {
	return &GlobalNode{Type: NodeGlobal}
}

func (l *GlobalNode) append(n Node) {
	l.Nodes = append(l.Nodes, n)
}

func (l *GlobalNode) String() string {
	b := new(bytes.Buffer)
	for _, n := range l.Nodes {
		fmt.Fprint(b, n)
	}
	return b.String()
}

func (l *GlobalNode) ToGolang(indent int) string {
	var funcs []string
	for _, fn := range l.Nodes {
		funcs = append(funcs, fn.ToGolang(0))
	}

	return fmt.Sprintf("%s\n", strings.Join(funcs, "\n\n"))
}
