package lexer

type SeparatorType string

const (
	SepLeftRoundBracket  SeparatorType = "("
	SepRightRoundBracket SeparatorType = ")"

	SepLeftSquareBracket  SeparatorType = "["
	SepRightSquareBracket SeparatorType = "]"

	SepLeftCurlyBracket  SeparatorType = "{"
	SepRightCurlyBracket SeparatorType = "}"

	SepComma     SeparatorType = ","
	SepSemicolon SeparatorType = ";"
)
