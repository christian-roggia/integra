package parser

import "fmt"

type CommentNode struct {
	Type    NodeType `json:"type"`
	Comment string   `json:"value"`
}

func newCommentNode(comment string) *CommentNode {
	return &CommentNode{Type: NodeComment, Comment: comment}
}

func (c *CommentNode) String() string {
	return c.Comment
}

func (c *CommentNode) ToGolang() string {
	return fmt.Sprintf("//%s", c.Comment[2:])
}
