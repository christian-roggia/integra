package parser

import (
	"fmt"
	"strings"
)

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

func (c *CommentNode) ToGolang(indent int) string {
	i := strings.Repeat(" ", indent*GolangIndent)
	return fmt.Sprintf("%s//%s", i, c.Comment[2:])
}
