package interpreter

import (
	"errors"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type Interpreter interface {
	Interpret(verb *ast.Verb) (ok bool, err error)
}

func InterpretBody(i Interpreter, node ast.Node) error {
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

		ok, err := i.Interpret(verb)
		if err != nil {
			return err
		}

		if !ok {
			return errors.New(fmt.Sprintf("Unknown config verb '%s'", verb.Key.Value))
		}
	}

	return nil
}

type Skip struct{}

func NewSkip() *Skip {
	return &Skip{}
}

func (s *Skip) Interpret(verb *ast.Verb) (bool, error) {
	return true, nil
}

type Pipeline struct {
	Interpreters []Interpreter
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) Use(i Interpreter) {
	p.Interpreters = append(p.Interpreters, i)
}

func (p *Pipeline) Interpret(verb *ast.Verb) (bool, error) {
	for _, i := range p.Interpreters {
		ok, err := i.Interpret(verb)
		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}

type Router struct {
	Route        func(verb *ast.Verb) string
	Interpreters map[string]Interpreter
}

func NewRouter(route func(verb *ast.Verb) string) *Router {
	return &Router{
		Route:        route,
		Interpreters: make(map[string]Interpreter),
	}
}

func (r *Router) Use(name string, i Interpreter) {
	pipeline := r.Get(name)

	if pipeline == nil {
		pipeline = NewPipeline()
		r.Interpreters[name] = pipeline
	}

	pipeline.Use(i)
}

func (r *Router) Get(name string) *Pipeline {
	return r.Interpreters[name].(*Pipeline)
}

func (r *Router) Interpret(verb *ast.Verb) (bool, error) {
	var route string
	if r.Route != nil {
		route = r.Route(verb)
	} else {
		route = verb.Key.Value
	}

	i, ok := r.Interpreters[route]
	if ok {
		return i.Interpret(verb)
	}

	return false, nil
}
