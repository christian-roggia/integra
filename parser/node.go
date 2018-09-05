package parser

// NodeType identifies the type of a parse tree node.
type NodeType string

const (
	NodeGlobal              NodeType = "GLOBAL"
	NodeCall                         = "CALL"
	NodeComment                      = "COMMENT"
	NodeEquation                     = "EQUATION"
	NodeExpression                   = "EXPRESSION"
	NodeIf                           = "IF"
	NodeLiteral                      = "LITERAL"
	NodeNumber                       = "NUMBER"
	NodeOperator                     = "OPERATOR"
	NodeReturn                       = "RETURN"
	NodeString                       = "STRING"
	NodeVariable                     = "VARIABLE"
	NodeVariableDeclarator           = "VARIABLE_DECLARATOR"
	NodeVariableInitializer          = "VARIABLE_INITIALIZER"
)

type Node interface {
	String() string
	ToGolang() string
}
