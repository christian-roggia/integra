package lexer

// KeywordType lists all the language specific keywords.
type KeywordType string

const (
	// KwEquation is used to define new functions.
	KwEquation KeywordType = "equation"

	// KwIf is used for if ... {} conditions.
	KwIf KeywordType = "if"

	// KwElse is used for if ... {} else {} conditions.
	KwElse KeywordType = "else"

	// KwReturn is used to return values from a function.
	KwReturn KeywordType = "return"

	// KwPrint is a built-in function that allows data to be printed in the console.
	KwPrint KeywordType = "print"

	// KwWrite is a built-in function that allows data to be saved in a file.
	KwWrite KeywordType = "write"

	// KwPi is a built-in mathematical function that will return the n-th decimal digit of pi.
	KwPi KeywordType = "if"

	// KwEuler is a built-in mathematical function that will return the n-th decimal digit of e.
	KwEuler KeywordType = "if"

	// KwPrime is a built-in mathematical function that will return the n-th prime number.
	KwPrime KeywordType = "if"
)

// Keywords is a collection of all the valid keywords.
var Keywords = []KeywordType{
	KwEquation,
	KwIf, KwElse, KwReturn,
	KwPrint, KwWrite,
	KwPi, KwEuler, KwPrime,
}

// IsKeyword returns whether the current string is a keyword or not.
func IsKeyword(s string) bool {
	for _, k := range Keywords {
		if s == string(k) {
			return true
		}
	}

	return false
}
