package permissions

import (
	"errors"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type Permissions struct {
	byname map[string]*ast.VerbBody
}

func New() *Permissions {
	return &Permissions{
		byname: make(map[string]*ast.VerbBody),
	}
}

func (p *Permissions) Register(name string, body *ast.VerbBody) {
	p.byname[name] = body
}

func (p *Permissions) Interpret(verb *ast.Verb) (bool, error) {
	// only match plain permission verb
	if verb.Key.Value != "permission" || len(verb.Key.Modifiers) != 0 {
		return false, nil
	}

	if verb.Values.IsEmpty() {
		return false, errors.New("Permission is missing a name.")
	}

	namenode := verb.Values.Items[0]
	name, ok := namenode.(*ast.MethodCall)
	if !ok || name.HasArguments() || !name.Path.IsEmpty() || name.Key.HasModifiers() {
		return false, errors.New("Expecting permission name.")
	}

	p.Register(name.Key.Value, verb.Body)
	return true, nil
}
