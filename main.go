package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/christian-roggia/integra/lexer"
	"github.com/christian-roggia/integra/parser"
)

func main() {
	generateLexerFile()

	b, _ := ioutil.ReadFile("integra.int")
	p := parser.NewParser("integra.int", string(b))
	if err := p.Parse(); err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	generateParserFile(p)
	generateCodeFile(p)
	generateSymbolFile(p)

	fmt.Println("Succesfully finished.")
}

func generateLexerFile() {
	b, _ := ioutil.ReadFile("integra.int")
	l := lexer.Lex("integra.int", string(b))

	s := ""
	for {
		select {
		case item := <-l.Get():
			s = fmt.Sprintf("%s%s\n", s, item.String())
			if item.Type == lexer.TokenEOF {
				ioutil.WriteFile("build/lexer.txt", []byte(s), 0644)
				return
			}
		}
	}
}

func generateParserFile(p *parser.Parser) {
	j, err := json.MarshalIndent(p.Tree(), "", "  ")
	if err != nil {
		fmt.Printf("json.Marshal(): %s\n", err)
		return
	}

	ioutil.WriteFile("build/parser.json", []byte(j), 0644)
}

func generateCodeFile(p *parser.Parser) {
	ioutil.WriteFile("build/main.go", []byte(p.Tree().ToGolang(0)), 0644)
	ioutil.WriteFile("build/_main.c", []byte(p.Tree().ToC(0)), 0644)
}

func generateSymbolFile(p *parser.Parser) {
	j, err := json.MarshalIndent(p.Symbols(), "", "  ")
	if err != nil {
		fmt.Printf("json.Marshal(): %s\n", err)
		return
	}

	ioutil.WriteFile("build/symbols.json", []byte(j), 0644)
}
