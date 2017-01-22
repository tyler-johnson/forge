package fcl

import (
	"io/ioutil"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
	"github.com/tyler-johnson/forge/forge/fcl/parser"
)

func Parse(src []byte) (*ast.File, error) {
	p := parser.New(src)
	return p.Parse()
}

func ParseFile(fn string) (*ast.File, error) {
	src, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	return Parse(src)
}
