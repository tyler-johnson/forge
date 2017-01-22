package ast

import "github.com/tyler-johnson/forge/forge/fcl/token"

// Node is an element in the abstract syntax tree.
type Node interface {
	node()
	Pos() token.Pos
}

func (File) node()       {}
func (Verb) node()       {}
func (VerbBody) node()   {}
func (Comment) node()    {}
func (Identifier) node() {}
func (Literal) node()    {}
func (ValueGroup) node() {}
func (Path) node()       {}
func (PathPart) node()   {}
func (Variable) node()   {}
func (MethodCall) node() {}

// File represents a single FCL file
type File struct {
	Node *VerbBody
}

// Pos is the starting position of the file
func (f *File) Pos() token.Pos {
	return f.Node.Pos()
}
