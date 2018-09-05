package lexer

type OperatorType string

const (
	OpGreater      = ">"
	OpLess         = "<"
	OpEqual        = "=="
	OpGreaterEqual = ">="
	OpLessEqual    = "<="

	OpAddition       = "+"
	OpSubstraction   = "-"
	OpMultiplication = "*"
	OpDivision       = "/"

	OpInitialization = ":="
	OpAssignment     = "="
)
