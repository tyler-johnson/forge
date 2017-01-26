package interpreter

import (
	"errors"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type VerbHandler func(*Context)

type Interpreter struct {
	verbs map[string]VerbHandler
}

func New() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) RegisterVerb(name string, fn VerbHandler) {
	i.verbs[name] = fn
}

func (i *Interpreter) Interpret(node ast.Node) error {
	var body *ast.VerbBody

	switch item := node.(type) {
	case *ast.File:
		body = item.Root
	case *ast.Verb:
		body = item.Body
	case *ast.VerbBody:
		body = item
	default:
		return errors.New("Expecting File, Verb or VerbBody to interpret.")
	}

	for _, child := range body.Items {
		verb, ok := child.(*ast.Verb)
		if !ok {
			continue
		}

		fmt.Println(verb.Key.Value)
	}

	return nil
}
