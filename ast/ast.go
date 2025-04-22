package ast

import (
	"bytes"
	"fmt"

	"monkey/token"
)

// All nodes must implement this interface
type Node interface {
	TokenLiteral() string
	String() string

	// required for %#v custom formatting
	GoString() string
}

// language constructs that perform actions but don't produce values
// embed: Node interface
type Statement interface {
	Node
	statementNode()
}

// language construct that produce values
// embed: Node interface
type Expression interface {
	Node
	expressionNode()
}

// Root Node of the AST
// Implements: Node
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p *Program) GoString() string {
	var out bytes.Buffer
	out.WriteString("&ast.Program{\n")
	for _, s := range p.Statements {
		out.WriteString(fmt.Sprintf("\t%#v,\n", s))
	}
	out.WriteString("}")
	return out.String()
}

// Represents a variable binding statement
// Format: let <identifier> = <expression>
// Implements: Statement
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // Variable name
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) GoString() string {
	return fmt.Sprintf("&ast.LetStatement{Token:%#v, Name:%#v, Value:%#v}",
		ls.Token, ls.Name, ls.Value)
}

// Implements: Expression
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // Value of the idenfitier, the name
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// Implement GoString for Identifier
func (i *Identifier) GoString() string {
	// Use %#v for the Token struct, %q for the string value
	return fmt.Sprintf("&ast.Identifier{Token:%#v, Value:%q}",
		i.Token, i.Value)
}

// Represents a return
// Format: return <expression>
// Implements: Statement
type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Implement GoString for ReturnStatement
func (rs *ReturnStatement) GoString() string {
	return fmt.Sprintf("&ast.ReturnStatement{Token:%#v, ReturnValue:%#v}",
		rs.Token, rs.ReturnValue)
}

// Represents a statement that consists solely of one expression.
// It's a wrapper
// e.g. x + 10;
// Implements Statement
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (es *ExpressionStatement) GoString() string {
	return fmt.Sprintf("&ast.ExpressionStatement{Token:%#v, Expression:%#v}",
		es.Token, es.Expression)
}
