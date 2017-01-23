package ast

import (
	"encoding/json"
	"strconv"

	"github.com/tyler-johnson/forge/forge/fcl/token"
)

type ValueGroup struct {
	Items []Node // Literal or Variable or MethodCall
}

func (a *ValueGroup) Pos() token.Pos {
	if a.IsEmpty() {
		return token.Pos{}
	}

	return a.Items[0].Pos()
}

func (a *ValueGroup) Add(n Node) {
	a.Items = append(a.Items, n)
}

func (a *ValueGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Items)
}

func (f *ValueGroup) NodeName() string {
	return "ValueGroup"
}

func (vg *ValueGroup) IsEmpty() bool {
	return vg.Items == nil || len(vg.Items) < 1
}

type Identifier struct {
	RawPos    token.Pos
	Modifiers []string
	Value     string
}

func (a *Identifier) Modifier(mod string) bool {
	for _, m := range a.Modifiers {
		if m == mod {
			return true
		}
	}

	return false
}

func (a *Identifier) SetModifier(mods ...string) {
	for _, mod := range mods {
		if !a.Modifier(mod) {
			a.Modifiers = append(a.Modifiers, mod)
		}
	}
}

func (f *Identifier) NodeName() string {
	return "Identifier"
}

func (a *Identifier) Pos() token.Pos {
	return a.RawPos
}

type LiteralType int

const (
	NULL LiteralType = iota
	STRING
	BOOL
	NUMBER
)

type Literal struct {
	RawPos token.Pos
	Type   LiteralType
	Value  interface{}
}

func (f *Literal) NodeName() string {
	return "Literal"
}

func (a *Literal) Pos() token.Pos {
	return a.RawPos
}

func (a *Literal) String() string {
	switch a.Type {
	case NULL:
		return "null"
	case STRING:
		v, _ := json.Marshal(a.Value)
		return string(v[:])
	case BOOL:
		return strconv.FormatBool(a.Value.(bool))
	case NUMBER:
		return strconv.FormatFloat(a.Value.(float64), 'f', -1, 64)
	default:
		return ""
	}
}

type Variable struct {
	Key  *Identifier
	Path *Path
}

func (f *Variable) NodeName() string {
	return "Variable"
}

func (a *Variable) Pos() token.Pos {
	return a.Key.Pos()
}

type MethodCall struct {
	Key       *Identifier
	Path      *Path
	Arguments *ValueGroup
}

func (f *MethodCall) NodeName() string {
	return "MethodCall"
}

func (a *MethodCall) Pos() token.Pos {
	return a.Key.Pos()
}
