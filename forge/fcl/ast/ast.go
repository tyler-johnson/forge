package ast

import "github.com/tyler-johnson/forge/forge/fcl/token"

// Node is an element in the abstract syntax tree.
type Node interface {
	Pos() token.Pos
	NodeName() string
}

// File represents a single FCL file
type File struct {
	Root *VerbBody
}

func (f *File) NodeName() string {
	return "File"
}

// Pos is the starting position of the file
func (f *File) Pos() token.Pos {
	return f.Root.Pos()
}
