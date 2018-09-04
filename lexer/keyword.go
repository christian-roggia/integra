package lexer

// KeywordType lists all the language specific keywords.
type KeywordType string

const (
	// KeywordEquation is used to define new functions.
	KeywordEquation KeywordType = "equation"

	// KeywordIf is used for if ... {} conditions.
	KeywordIf KeywordType = "if"

	// KeywordElse is used for if ... {} else {} conditions.
	KeywordElse KeywordType = "else"

	// KeywordReturn is used to return values from a function.
	KeywordReturn KeywordType = "return"

	// KeywordPrint is a built-in function that allows data to be printed in the console.
	KeywordPrint KeywordType = "print"

	// KeywordWrite is a built-in function that allows data to be saved in a file.
	KeywordWrite KeywordType = "write"

	// KeywordPi is a built-in mathematical function that will return the n-th decimal digit of pi.
	KeywordPi KeywordType = "if"

	// KeywordEuler is a built-in mathematical function that will return the n-th decimal digit of e.
	KeywordEuler KeywordType = "if"

	// KeywordPrime is a built-in mathematical function that will return the n-th prime number.
	KeywordPrime KeywordType = "if"
)

// Keywords is a collection of all the valid keywords.
var Keywords = []KeywordType{
	KeywordEquation,
	KeywordIf, KeywordElse, KeywordReturn,
	KeywordPrint, KeywordWrite,
	KeywordPi, KeywordEuler, KeywordPrime,
}
