# Welcome to Integra
Integra is a programming language designed for pure mathematics. This programming language has been design as open source project for the **Free University of Bolzano**.

## Lexer
The lexer of Integra uses goroutines and channels to feed tokens to the parser. This is a very common approach used in Golang to allow concurrent operations and therefore have high performances compared to traditional Lexer and Parser that will execute lexical analysis and syntatical analysis sequantially.

## Parser
The Parser will generate an AST and produce a Symbol Table. The parser will consume data coming from the Lexer as soon as it is ready using concurrency. This is a recursive descend parser.

## Abstract Syntax Tree

## Symbol Table

## Code Generation