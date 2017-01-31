// Copyright 2016 Marapongo, Inc. All rights reserved.

package symbols

import (
	"github.com/marapongo/mu/pkg/compiler/ast"
	"github.com/marapongo/mu/pkg/diag"
	"github.com/marapongo/mu/pkg/tokens"
)

// Variable is an interface common to all variables.
type Variable interface {
	Symbol
	Type() Type
	VarNode() ast.Variable
}

var _ Symbol = (Variable)(nil)

// LocalVariable is a fully bound local variable symbol.
type LocalVariable struct {
	Node *ast.LocalVariable
	Nm   tokens.Name
	Ty   Type
}

var _ Symbol = (*LocalVariable)(nil)
var _ Variable = (*LocalVariable)(nil)

func (node *LocalVariable) symbol()               {}
func (node *LocalVariable) Name() tokens.Name     { return node.Nm }
func (node *LocalVariable) Token() tokens.Token   { return tokens.Token(node.Name()) }
func (node *LocalVariable) Tree() diag.Diagable   { return node.Node }
func (node *LocalVariable) Type() Type            { return node.Ty }
func (node *LocalVariable) VarNode() ast.Variable { return node.Node }
func (node *LocalVariable) String() string        { return string(node.Name()) }

// NewLocalVariableSym returns a new LocalVariable symbol associated with the given AST node.
func NewLocalVariableSym(node *ast.LocalVariable, ty Type) *LocalVariable {
	return &LocalVariable{Node: node, Nm: node.Name.Ident, Ty: ty}
}

// NewSpecialVariableSym returns a "special" LocalVariable symbol that has no corresponding AST node and has a name.
func NewSpecialVariableSym(nm tokens.Name, ty Type) *LocalVariable {
	return &LocalVariable{Node: nil, Nm: nm, Ty: ty}
}