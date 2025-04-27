package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"mi-proyecto/lexer"
	"mi-proyecto/token"
)

// EOF_TOKEN se usa para comparar con el token de fin de entrada.
var EOF_TOKEN = token.Token{Type: token.EOF, Literal: ""}

// startREPL inicia el bucle de lectura-evaluación-impresión.
func startREPL() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error leyendo entrada:", err)
			return
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}

		l := lexer.New(input)
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Println(tok)
		}
	}
}

func main() {
	startREPL()
}
