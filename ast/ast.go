package ast

import (
	"strings"
	"mi-proyecto/token"
)

// Node es la interfaz base para todos los nodos del AST.
type Node interface {
	// TokenLiteral retorna el literal del token asociado al nodo.
	TokenLiteral() string
	// String retorna una representación legible del nodo.
	String() string
}

// Statement es un nodo que representa una declaración.
type Statement interface {
	Node
	statementNode()
}

// Expression es un nodo que representa una expresión.
type Expression interface {
	Node
	expressionNode()
}

// Program es el nodo raíz del AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral retorna el literal del primer token en el programa.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String concatena la representación de cada statement.
func (p *Program) String() string {
	var out strings.Builder
	for _, s := range p.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

// LetStatement representa la declaración 'hatsune'.
type LetStatement struct {
	Token token.Token // El token 'hatsune'
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral retorna el literal del token 'hatsune'.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String formatea la declaración hatsune: 'hatsune <nombre> = <valor>'.
func (ls *LetStatement) String() string {
	var out strings.Builder
	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	return out.String()
}

// Identifier representa un identificador en el AST.
type Identifier struct {
	Token token.Token // El token del identificador
	Value string      // El nombre literal
}

func (i *Identifier) expressionNode() {}

// TokenLiteral retorna el literal del token de identificador.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String retorna el valor del identificador.
func (i *Identifier) String() string {
	return i.Value
}
