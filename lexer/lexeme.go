package lexer

import (
	"fmt"
	"text/scanner"
	"unicode"
)

func lexText(l *Lexer) stateFn {
	for r := l.next(); isWhitespace(r); l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == scanner.EOF:
		l.emit(TokenEOF)
		return nil
	case isSeparator(r, l.peek()):
		l.backup()
		return lexSeparator
	case r == '"':
		return lexString
	case isNumber(r, l.peek()):
		l.backup()
		return lexNumber
	case isComment(r, l.peek()):
		return lexComment
	case isIdentifier(r):
		return lexIdentifier
	case isOperator(r, l.peek()):
		l.backup()
		return lexOperator
	default:
		l.errorf(fmt.Sprintf("don't know what to do with: %q", r))
		return nil
	}
}

func lexComment(l *Lexer) stateFn {
	for {
		if l.peek() == '\n' && l.peek() != scanner.EOF {
			break
		}
		l.next()
	}

	l.emit(TokenComment)
	return lexText
}

func lexIdentifier(l *Lexer) stateFn {
	allow := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	l.acceptRun(allow)

	for _, k := range Keywords {
		if l.input[l.start:l.pos] == string(k) {
			l.emit(TokenKeyword)
			return lexText
		}
	}

	l.emit(TokenIdentifier)
	return lexText
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != '"'; r = l.next() {
		if r == '\r' {
			r = l.next()
		}
		if r == '\n' || r == scanner.EOF {
			return l.errorf("unterminated quoted string")
		}
	}

	l.emit(TokenString)
	return lexText
}

func lexNumber(l *Lexer) stateFn {
	l.accept("+-")

	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}

	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}

	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}

	l.emit(TokenNumber)
	return lexText
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

func isNumber(r, n rune) bool {
	if r == '+' || r == '-' {
		return ('0' <= n && n <= '9')
	}

	return ('0' <= r && r <= '9')
}

func isIdentifier(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isComment(r, n rune) bool {
	return r == '/' && n == '/'
}

func isSeparator(r, n rune) bool {
	return r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' || r == ',' || r == ';'
}

func isOperator(r, n rune) bool {
	return r == '>' || r == '<' || r == '=' || r == '+' || r == '-' || r == '*' || r == '/' || r == ':'
}

func lexSeparator(l *Lexer) stateFn {
	l.accept("()[]{},;")
	l.emit(TokenSeparator)
	return lexText
}

func lexOperator(l *Lexer) stateFn {
	if l.accept(":") && !l.accept("=") {
		l.errorf("unexpected colon")
	}

	if l.accept("<>=") {
		l.accept("=")
	}

	l.accept("<>=*/+-")
	l.emit(TokenOperator)
	return lexText
}
