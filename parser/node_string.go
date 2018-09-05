package parser

type StringNode struct {
	Type  NodeType `json:"type"`
	Value string   `json:"value"`
}

func newStringNode(v string) *StringNode {
	return &StringNode{Type: NodeString, Value: v}
}

func (n *StringNode) String() string {
	return n.Value
}

func (n *StringNode) ToGolang(indent int) string {
	return n.Value
}
