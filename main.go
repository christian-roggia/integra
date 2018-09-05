package main

import (
	"fmt"
	"io/ioutil"

	"github.com/christian-roggia/integra/parser"
)

func main() {
	/*b, _ := ioutil.ReadFile("integra.int")
	l := lexer.Lex("integra.int", string(b))
	for {
		select {
		case item := <-l.Get():
			fmt.Printf("%s\n", item.String())
			if item.Type == lexer.TokenEOF {
				return
			}
		}
	}*/

	b, _ := ioutil.ReadFile("integra.int")
	p := parser.NewParser("integra.int", string(b))
	if err := p.Parse(); err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	// j, err := json.MarshalIndent(p.Tree(), "", "  ")
	// if err != nil {
	// 	fmt.Printf("json.Marshal(): %s\n", err)
	// 	return
	// }

	// fmt.Printf("%s\n", j)

	fmt.Printf("%s\n", p.Tree().ToGolang())
	fmt.Println("Succesfully finished.")
}
