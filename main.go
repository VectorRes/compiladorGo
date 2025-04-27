package main

import (
	"fmt"
	"io/ioutil"

	"mi-proyecto/lexer"
	"mi-proyecto/token"
)

func runFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Archivo no encontrado: %s\n", filename)
		return
	}

	l := lexer.New(string(data))
	
	// Iterar sobre los tokens hasta encontrar EOF
	for {
		tok := l.NextToken()
		fmt.Printf("Type: %s, Literal: %s\n", tok.Type, tok.Literal)
		
		if tok.Type == token.EOF {
			break
		}
	}
}

func main() {
	runFile("compilador.miku")
}