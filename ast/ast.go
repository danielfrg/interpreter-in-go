package ast

import "monkey/token"

// All nodes must implement this interface
type Node interface {
	TokenLiteral() string
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

// Implements: Expression
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // Value of the idenfitier, the name
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
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
