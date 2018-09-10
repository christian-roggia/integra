package parser

// This is a semplified C grammar.
// https://www.lysator.liu.se/c/ANSI-C-grammar-y.html

// function_definition
// 		EQUATION declarator compound_statement

// declarator
// 		: IDENTIFIER '(' identifier_list ')'
// 		| IDENTIFIER '(' ')'
// 		;

// identifier_list
// 		: IDENTIFIER
// 		| identifier_list ',' IDENTIFIER
// 		;

// compound_statement
//		: '{' '}'
//		| '{' statement_list '}'
//		;

// statement_list
// 		: statement
// 		| statement_list statement
// 		;

// statement
// 	: compound_statement
// 	| expression_statement
// 	| selection_statement
// 	| jump_statement
//  | init_declarator
// 	;

// expression_statement
// 	: ';'
// 	| expression ';'
// 	;

// selection_statement
// 	: IF '(' expression ')' statement
// 	;

// jump_statement
// 	| RETURN expression ';'
// 	;

// expression
// 	: assignment_expression
// 	| expression ',' assignment_expression
// 	;

// init_declarator
// 	: IDENTIFIER ':' '=' initializer ';'
// 	;
