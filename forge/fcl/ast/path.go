package ast

import (
	"encoding/json"

	"github.com/tyler-johnson/forge/forge/fcl/token"
)

type Path struct {
	Parts []*PathPart
}

func (a *Path) Pos() token.Pos {
	if a.IsEmpty() {
		return token.Pos{}
	}

	return a.Parts[0].Pos()
}

func (a *Path) Add(n *PathPart) {
	a.Parts = append(a.Parts, n)
}

func (a *Path) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Parts)
}

func (f *Path) NodeName() string {
	return "Path"
}

func (vg *Path) IsEmpty() bool {
	return vg.Parts == nil || len(vg.Parts) < 1
}

type PathPartType int

const (
	SIMPLE PathPartType = iota
	COMPOUND
)

type PathPart struct {
	RawPos token.Pos
	Type   PathPartType
	Value  interface{}
}

func (a *PathPart) Pos() token.Pos {
	return a.RawPos
}

func (f *PathPart) NodeName() string {
	return "PathPart"
}
