package models

import (
	"errors"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
	"github.com/tyler-johnson/forge/forge/fcl/printer"
)

type Models struct {
	ByName map[string]*Model
}

func New() *Models {
	return &Models{make(map[string]*Model)}
}

func (m *Models) Get(name string) *Model {
	model, ok := m.ByName[name]
	if !ok {
		model = &Model{}
		m.ByName[name] = model
	}

	return model
}

func (m *Models) Interpret(verb *ast.Verb) (bool, error) {
	// only match plain model verb
	if verb.Key.Value != "model" || len(verb.Key.Modifiers) != 0 {
		return false, nil
	}

	model := m.Get(verb.Key.Value)
	err := model.Interpret(verb.Body)
	return true, err
}

type Model struct {
	Schema *Schema
}

func NewModel() *Model {
	return &Model{&Schema{}}
}

func (m *Model) Interpret(vb *ast.VerbBody) error {
	// parse children
	return nil
}

type Schema struct {
}

func (s *Schema) Interpret(verb *ast.Verb) (bool, error) {
	// look for schema declaration
	return false, nil
}

type SchemaField struct {
	Key    string
	Type   string
	Traits []string
}

func (sf *SchemaField) Interpret(verb *ast.Verb) (bool, error) {
	if len(verb.Key.Modifiers) != 0 {
		return false, nil
	}

	sf.Key = verb.Key.Value

	if !verb.Values.IsEmpty() {
		typedec, ok := verb.Values.Items[0].(*ast.MethodCall)
		if !ok {
			return false, errors.New(fmt.Sprintf("Expecting type name as method, instead got '%s'", verb.Values.Items[0]))
		}

		typebyte, err := printer.Print(typedec)
		if err != nil {
			return false, err
		}

		sf.Type = string(typebyte[:])
	}

	return true, nil
}
