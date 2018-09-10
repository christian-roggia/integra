# Welcome to Integra
Integra is a programming language designed for pure mathematics. This programming language has been design as open source project for the **Free University of Bolzano**. This repository provides functionalities for transpiling (Source-to-Source compiling) from Integra to Golang and C source code and provides an interface for the easy implementation of other programming language. This compiler includes full lexical analysis and syntatical analysis.

## Lexer
The lexer of Integra uses goroutines and channels to feed tokens to the parser. This is a very common approach used in Golang to allow concurrent operations and therefore have high performances compared to traditional Lexers and Parsers that will execute lexical analysis and syntatical analysis sequantially.

## Parser
The Parser will generate an AST and produce a Symbol Table. The parser will consume data coming from the Lexer as soon as it is ready using the concurrent goroutine. This is a recursive descent parser of type LL(k). Integra has no variable typing as the only type allowed is integer (64bit). The only exception where strings are allowed is only for calling the built-in functions "print" (console logging function) and "write" (persistent file writing function), which are considered reserved programming keywords.  

A grammar of the Parser has been described using YACC notation and is a mixed grammar of the C ang Golang grammars. This grammar is only a representation as the real grammar is implemented at source level without using third parties libraries or frameworks (i.e. LEX and YACC).  

All functions, with the exception of "write" and "print", must return a value and only a single value. It is also not allowed to call functions outside an expression or an assignment, as all functions are expected to return a value, again with the exception of "write" and "print".

## Abstract Syntax Tree
The representation of the Abstract Syntax Tree is generated using a node tree and the output is stored in JSON format for an easy visualization.

## Symbol Table
The symbol table generated checks for duplicated initialization, variable shadowing and variable assignment without variable initialization. It also verify that all variables used in expressions have previously been declared.

## Code Generation
The code generation does not generate native code. The transpiler is instead capable of generating source code for Golang, C and potentially other programming languages.