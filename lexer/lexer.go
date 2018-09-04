package lexer

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode/utf8"
)

type stateFn func(*Lexer) stateFn

// Lexer represents the lexer for our programming language.
type Lexer struct {
	name  string
	input string
	state stateFn
	pos   int
	start int
	width int

	tokens chan Token
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return scanner.EOF
	}

	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width

	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// emit passes a Token back to the client.
func (l *Lexer) emit(t TokenType) {
	l.tokens <- Token{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore will ignore the current token.
func (l *Lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// errorf interrupts the lexing process and emits an error.
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- Token{TokenError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

// Get returns the tokens channel.
func (l *Lexer) Get() chan Token {
	return l.tokens
}

// Lex will analyze all the text passed via the input.
func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:   name,
		input:  input,
		tokens: make(chan Token),
	}

	go l.run()
	return l
}

// run consumes the next token until EOF or an error is encountered.
func (l *Lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.tokens)
}
