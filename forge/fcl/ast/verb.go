package ast

import (
	"encoding/json"

	"github.com/tyler-johnson/forge/forge/fcl/token"
)

type VerbBody struct {
	Items    []Node // Comment or Verb
	StartPos token.Pos
	EndPos   token.Pos
}

func (f *VerbBody) NodeName() string {
	return "VerbBody"
}

func (vb *VerbBody) Pos() token.Pos {
	return vb.StartPos
}

func (vb *VerbBody) Add(n Node) {
	vb.Items = append(vb.Items, n)
}

func (vb *VerbBody) IsEmpty() bool {
	return vb.Items == nil || len(vb.Items) < 1
}

func (a *VerbBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Items)
}

// Verb is the representation of high-level action within Forge.
type Verb struct {
	Key    *Identifier
	Values *ValueGroup
	Body   *VerbBody
}

func (f *Verb) NodeName() string {
	return "Verb"
}

func (v *Verb) Pos() token.Pos {
	return v.Key.RawPos
}
