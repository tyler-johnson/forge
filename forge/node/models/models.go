package models

import (
	"errors"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
	"github.com/tyler-johnson/forge/forge/fcl/printer"
	"github.com/tyler-johnson/forge/forge/node/interpreter"
	"github.com/tyler-johnson/forge/forge/node/utils"
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
		model = NewModel(name)
		m.ByName[name] = model
	}

	return model
}

func (m *Models) Interpret(verb *ast.Verb) (bool, error) {
	// only match plain model verb
	if !utils.MatchKey(verb, "model", nil) {
		return false, nil
	}

	name, ok := utils.ExtractName(verb)
	if !ok {
		return false, errors.New(fmt.Sprintf("Expecting type name as method, instead got '%s'", verb.Values.Items[0]))
	}

	return m.Get(name).Interpret(verb)
}

type Model struct {
	Name   string
	Schema *Schema
	pipe   *interpreter.Pipeline
}

func NewModel(name string) *Model {
	schema := NewSchema()
	pipe := interpreter.NewPipeline()
	pipe.Use(schema)
	pipe.Use(interpreter.NewSkip())

	return &Model{
		Name:   name,
		Schema: schema,
		pipe:   pipe,
	}
}

func (m *Model) Interpret(verb *ast.Verb) (bool, error) {
	err := interpreter.InterpretBody(m.pipe, verb)
	return true, err
}

type Schema struct {
	ByKey map[string]*SchemaField
}

func NewSchema() *Schema {
	return &Schema{
		ByKey: make(map[string]*SchemaField),
	}
}

func (s *Schema) Interpret(verb *ast.Verb) (bool, error) {
	if !utils.MatchKey(verb, "schema", nil) {
		return false, nil
	}

	for _, node := range verb.Body.Items {
		switch item := node.(type) {
		case *ast.Verb:
			if verb.Key.HasModifiers() {
				continue
			}

			field := &SchemaField{}
			field.Key = item.Key.Value
			s.ByKey[field.Key] = field
		}
	}

	return true, nil
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
