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
// 	;

// expression_statement
// 	: ';'
// 	| expression ';'
// 	;

// selection_statement
// 	: IF '(' expression ')' statement
// 	| IF '(' expression ')' statement ELSE statement
// 	;

// jump_statement
// 	| RETURN expression ';'
// 	;

// expression
// 	: assignment_expression
// 	| expression ',' assignment_expression
// 	;

// compound_statement
//		: '{' '}'
//		| '{' statement_list '}'
//		| '{' declaration_list '}'
//		| '{' declaration_list statement_list '}'
//		;

// declaration_list
//		: init_declarator
//		;

// init_declarator
// 	: IDENTIFIER ':' '=' initializer ';'
// 	;
