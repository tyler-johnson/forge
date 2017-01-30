package permissions

import (
	"errors"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type Permissions struct {
	ByName map[string]*Permission
}

func New() *Permissions {
	return &Permissions{
		ByName: make(map[string]*Permission),
	}
}

func (p *Permissions) Register(name string) {
	p.ByName[name] = &Permission{name}
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

	p.Register(name.Key.Value)
	return true, nil
}

type Permission struct {
	Name string
}
