package token

import (
	"fmt"
)

// TokenType representa el tipo de un token en el intérprete.
type TokenType int

// Definición de los tipos de token. Los valores son únicos y definidos via iota.
const (
	ASSIGN TokenType = iota + 1 // =
	COMMA                       // ,
	EOF                         // fin de archivo
	EQ                          // ==
	FOR                         // for
	FUNCTION                    // function
	IDENT                       // identificadores
	ILLEGAL                     // token ilegal o desconocido
	INT                         // enteros
	LBRACE                      // {
	HATSUNE                     // hatsune (forma de declarar variables)
	LPAREN                      // (
	PLUS                        // +
	RBRACE                      // }
	RPAREN                      // )
	SEMICOLON                   // ;
)

// String implementa fmt.Stringer para TokenType.
func (t TokenType) String() string {
	switch t {
	case ASSIGN:
		return "ASSIGN"
	case COMMA:
		return "COMMA"
	case EOF:
		return "EOF"
	case EQ:
		return "EQ"
	case FOR:
		return "FOR"
	case FUNCTION:
		return "FUNCTION"
	case IDENT:
		return "IDENT"
	case ILLEGAL:
		return "ILLEGAL"
	case INT:
		return "INT"
	case LBRACE:
		return "LBRACE"
	case HATSUNE:
		return "HATSUNE"
	case LPAREN:
		return "LPAREN"
	case PLUS:
		return "PLUS"
	case RBRACE:
		return "RBRACE"
	case RPAREN:
		return "RPAREN"
	case SEMICOLON:
		return "SEMICOLON"
	default:
		return "UNKNOWN"
	}
}

// Token representa un token con su tipo y literal.
type Token struct {
	Type    TokenType
	Literal string
}

// String implementa fmt.Stringer para Token.
func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %q)", t.Type, t.Literal)
}

// mapa de palabras clave que asocia literales con su TokenType.
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"hatsune":  HATSUNE,
	"for":      FOR,
}

// LookupTokenType devuelve el TokenType correspondiente a un literal:
// si es palabra clave, retorna su tipo; de lo contrario IDENT.
func LookupTokenType(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
