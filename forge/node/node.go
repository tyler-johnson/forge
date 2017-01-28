package node

import (
	"github.com/tyler-johnson/forge/forge/fcl/ast"
	"github.com/tyler-johnson/forge/forge/node/interpreter"
	"github.com/tyler-johnson/forge/forge/node/models"
	"github.com/tyler-johnson/forge/forge/node/permissions"
)

type Node struct {
	Perms  *permissions.Permissions
	Models *models.Models
}

func New() *Node {
	return &Node{
		Perms:  permissions.New(),
		Models: models.New(),
	}
}

func (n *Node) Interpret(tree *ast.File) error {
	pipeline := interpreter.NewPipeline()
	pipeline.Use(n.Perms)
	pipeline.Use(n.Models)
	pipeline.Use(interpreter.NewSkip())
	return interpreter.InterpretBody(pipeline, tree)
}
