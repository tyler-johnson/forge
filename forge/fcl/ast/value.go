package ast

import "github.com/tyler-johnson/forge/forge/fcl/token"

type ValueGroup struct {
	Items []Node // Literal or Variable or MethodCall
}

func (a *ValueGroup) Pos() token.Pos {
	return a.Items[0].Pos()
}

func (a *ValueGroup) Add(n Node) {
	a.Items = append(a.Items, n)
}

type Identifier struct {
	pos       token.Pos
	Modifiers []string
	Value     string
}

func NewIdentifier(pos token.Pos) *Identifier {
	return &Identifier{pos: pos}
}

func (a *Identifier) Pos() token.Pos {
	return a.pos
}

type LiteralType int

const (
	NULL LiteralType = iota
	STRING
	BOOL
	NUMBER
)

type Literal struct {
	pos   token.Pos
	Type  LiteralType
	Value interface{}
}

func NewLiteral(pos token.Pos) *Literal {
	return &Literal{pos: pos}
}

func (a *Literal) Pos() token.Pos {
	return a.pos
}

type Path struct {
	Parts []*PathPart
}

func NewPath() *Path {
	return &Path{make([]*PathPart, 0)}
}

func (a *Path) Pos() token.Pos {
	return a.Parts[0].Pos()
}

func (a *Path) Add(n *PathPart) {
	a.Parts = append(a.Parts, n)
}

type PathPartType int

const (
	SIMPLE PathPartType = iota
	COMPOUND
)

type PathPart struct {
	pos   token.Pos
	Type  PathPartType
	Value interface{}
}

func NewPathPart(pos token.Pos) *PathPart {
	return &PathPart{pos: pos}
}

func (a *PathPart) Pos() token.Pos {
	return a.pos
}

type Variable struct {
	Key  *Identifier
	Path *Path
}

func (a *Variable) Pos() token.Pos {
	return a.Key.Pos()
}

type MethodCall struct {
	Key       *Identifier
	Path      *Path
	Arguments *ValueGroup
}

func (a *MethodCall) Pos() token.Pos {
	return a.Key.Pos()
}
