package printer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type Printer struct {
	writer io.Writer
	line   int
	tablen int
}

func Print(node ast.Node) ([]byte, error) {
	p := New()
	return p.Print(node)
}

func New() *Printer {
	return &Printer{}
}

func (p *Printer) Print(node ast.Node) (b []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(runtime.Error); ok {
				panic(r)
			}
			if s, ok := r.(string); ok {
				panic(s)
			}
			err = r.(error)
		}
	}()

	p.reset()
	buf := bytes.NewBuffer(nil)
	defer func() {
		buf.Reset()
		buf = nil
	}()
	p.writer = buf

	switch i := node.(type) {
	case *ast.File:
		p.verbBody(i.Root)
	case *ast.VerbBody:
		p.verbBody(i)
	case *ast.Verb:
		p.verb(i)
	case *ast.Comment:
		p.comment(i)
	case *ast.Identifier:
		p.identifier(i)
	case *ast.ValueGroup:
		p.valueGroup(i, nil)
	case *ast.Literal:
		p.literal(i)
	case *ast.MethodCall:
		p.methodCall(i)
	case *ast.Variable:
		p.variable(i)
	case *ast.Path:
		p.path(i)
	default:
		return nil, errors.New(fmt.Sprintf("Unknown node '%s'", node.NodeName()))
	}

	b = make([]byte, buf.Len())
	copy(b, buf.Bytes())

	return
}

func (p *Printer) reset() {
	p.writer = nil
	p.line = 1
	p.tablen = 0
}

func (p *Printer) write(str string) int {
	n, err := p.writer.Write([]byte(str))
	if err != nil {
		panic(err)
	}

	return n
}

func (p *Printer) indent() int {
	p.tablen++
	return p.tablen
}

func (p *Printer) outdent() int {
	len := p.tablen - 1
	if len < 0 {
		len = 0
	}
	p.tablen = len
	return len
}

func (p *Printer) tabs() int {
	return p.write(strings.Repeat("  ", p.tablen))
}

func (p *Printer) newline() int {
	return p.write("\n")
}

func (p *Printer) space() int {
	return p.write(" ")
}

func (p *Printer) verbBody(body *ast.VerbBody) (n int) {
	if body.IsEmpty() {
		return
	}

	var prevline int

	for _, item := range body.Items {
		line := item.Pos().Line

		if prevline > 0 {
			if line != prevline {
				n += p.newline()
			}

			if line-prevline > 1 {
				n += p.newline()
			}
		}

		switch node := item.(type) {
		case *ast.Comment:
			if line == prevline {
				n += p.space()
			} else {
				n += p.tabs()
			}

			n += p.comment(node)
		case *ast.Verb:
			n += p.tabs() + p.verb(node)
		}

		prevline = line
	}

	n += p.newline()
	return
}

func (p *Printer) comment(c *ast.Comment) int {
	return p.write(c.Text)
}

func (p *Printer) verb(verb *ast.Verb) (n int) {
	n += p.identifier(verb.Key)
	n += p.valueGroup(verb.Values, func(n int) string {
		return " "
	})

	if verb.HasBody() {
		if !verb.Body.IsEmpty() {
			n += p.write(" {\n")
			p.indent()
			n += p.verbBody(verb.Body)
			p.outdent()
			n += p.tabs() + p.write("}")
		} else {
			n += p.write(" {}")
		}
	}

	return
}

func (p *Printer) identifier(id *ast.Identifier) (n int) {
	return p.modifiers(id.Modifiers) + p.write(id.Value)
}

func (p *Printer) valueGroup(vg *ast.ValueGroup, pre func(int) string) (n int) {
	if vg.IsEmpty() {
		return
	}

	for i, item := range vg.Items {
		if pre != nil {
			n += p.write(pre(i))
		} else if i != 0 {
			n += p.write(" ")
		}

		switch node := item.(type) {
		case *ast.Literal:
			n += p.literal(node)
		case *ast.MethodCall:
			n += p.methodCall(node)
		case *ast.Variable:
			n += p.variable(node)
		}
	}

	return
}

func (p *Printer) literal(lit *ast.Literal) (n int) {
	return p.write(lit.String())
}

func (p *Printer) methodCall(met *ast.MethodCall) (n int) {
	n += p.identifier(met.Key) + p.path(met.Path)
	if met.HasArguments() {
		n += p.write("(")
		n += p.valueGroup(met.Arguments, func(n int) string {
			if n > 0 {
				return " "
			}

			return ""
		})
		n += p.write(")")
	}
	return
}

func (p *Printer) variable(v *ast.Variable) (n int) {
	n += p.modifiers(v.Key.Modifiers)
	n += p.write("$") + p.write(v.Key.Value)
	n += p.path(v.Path)
	return
}

func (p *Printer) modifiers(mods []string) int {
	return p.write(strings.Join(mods, ""))
}

func (p *Printer) path(path *ast.Path) (n int) {
	for _, part := range path.Parts {
		switch part.Type {
		case ast.SIMPLE:
			n += p.write(".") + p.write(part.Value.(string))
		}
	}

	return
}
