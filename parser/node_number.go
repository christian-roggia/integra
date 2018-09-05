package parser

type NumberNode struct {
	Type  NodeType `json:"type"`
	Value string   `json:"value"`
}

func newNumberNode(v string) *NumberNode {
	return &NumberNode{Type: NodeNumber, Value: v}
}

func (n *NumberNode) String() string {
	return n.Value
}

func (n *NumberNode) ToGolang() string {
	return n.Value
}
