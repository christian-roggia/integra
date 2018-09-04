package main

import (
	"fmt"
	"io/ioutil"

	"github.com/christian-roggia/integra/lexer"
)

func main() {
	b, _ := ioutil.ReadFile("integra.int")
	l := lexer.Lex("integra.int", string(b))
	for {
		select {
		case item := <-l.Get():
			fmt.Printf("%s\n", item.String())
			if item.Type == lexer.TokenEOF {
				return
			}
		}
	}
}
