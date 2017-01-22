package ast

import "github.com/tyler-johnson/forge/forge/fcl/token"

type VerbBody struct {
	Items []Node // Comment or Verb
}

func (vb *VerbBody) Pos() token.Pos {
	return vb.Items[0].Pos()
}

func (vb *VerbBody) Add(n Node) {
	vb.Items = append(vb.Items, n)
}

// Verb is the representation of high-level action within Forge.
type Verb struct {
	Key    *Identifier
	Values *ValueGroup
	Body   *VerbBody
}

func NewVerb() *Verb {
	return &Verb{nil, &ValueGroup{}, &VerbBody{}}
}

func (v *Verb) Pos() token.Pos {
	return v.Key.pos
}
