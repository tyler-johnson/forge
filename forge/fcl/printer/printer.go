package printer

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
)

type Printer struct {
	file   *ast.File
	writer io.Writer
	buffer *bytes.Buffer
	line   int
	tablen int
}

func Print(file *ast.File) ([]byte, error) {
	p := New(file)
	return p.Print()
}

func New(file *ast.File) *Printer {
	return &Printer{file: file}
}

func (p *Printer) Print() (b []byte, err error) {
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
	p.buffer = bytes.NewBuffer(nil)
	defer func() {
		p.buffer.Reset()
		p.buffer = nil
	}()
	p.writer = p.buffer

	p.verbBody(p.file.Root)
	b = make([]byte, p.buffer.Len())
	copy(b, p.buffer.Bytes())

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

// func columnSizes(nodes []ast.Node, ct int) []int {
// 	sizes := make([]int, ct + 2)

// 	for _, item := range nodes {
// 		switch node := item.(type) {
// 		case *ast.Verb:
// 			n := len(node.Key.Modifiers) + len(node.Key.Value)
// 			if n > sizes[0] {
// 				sizes[0] = n
// 			}

//       if !node.Values.IsEmpty() {
//         for i := 0; i < ct; i++ {
//           if len(node.Values) > i {
//             sizes[i + 1] = len(node.Values[i].)
//           }
//         }
//       }
// 		}
// 	}

// 	return sizes
// }

type verbGroupColumns struct {
	items   []ast.Node
	values  int
	braces  bool
	comment bool
}

func parseColumns(nodes []ast.Node, valueCount int) (out []*verbGroupColumns) {
	type lineitem struct {
		items   []ast.Node
		values  int
		comment bool
		braces  bool
	}

	type group struct {
		lineitems []*lineitem
	}

	var prevline int
	var line *lineitem
	groups := make([]*group, 0)

	lastgroup := func() *group {
		len := len(groups)
		if len < 1 {
			return nil
		}
		return groups[len-1]
	}

	newgroup := func() *group {
		if last := lastgroup(); last != nil && len(last.lineitems) == 0 {
			return last
		}

		gr := &group{}
		groups = append(groups, gr)
		return gr
	}

	for _, item := range nodes {
		lineno := item.Pos().Line
		if line == nil || lineno != prevline {
			gr := lastgroup()
			if gr == nil {
				gr = newgroup()
			}

			if line != nil {
				gr.lineitems = append(gr.lineitems, line)
			}

			if lineno-prevline > 1 {
				gr = newgroup()
			}

			line = &lineitem{}
		}

		line.items = append(line.items, item)

		switch node := item.(type) {
		case *ast.Comment:
			line.comment = true
		case *ast.Verb:
			if node.Body != nil {
				line.braces = true
			}

			if node.Values != nil {
				line.values = len(node.Values.Items)
			}
		}

		prevline = lineno
	}

	if gr := lastgroup(); line != nil && gr != nil {
		gr.lineitems = append(gr.lineitems, line)
	}

	for _, group := range groups {
		fmt.Println(group)

	}

	return
}

func (p *Printer) verbBody(body *ast.VerbBody) (n int) {
	if body.IsEmpty() {
		return
	}

	groups := parseColumns(body.Items, 1)
	fmt.Println(groups)

	return

	// buffer := p.writer
	// var prevline int
	// var tabs *tabwriter.Writer

	// flush := func() {
	// 	if tabs != nil {
	// 		tabs.Flush()
	// 		p.writer = buffer
	// 		tabs = nil
	// 	}
	// }

	// for _, item := range body.Items {
	// 	line := item.Pos().Line

	// 	if prevline > 0 {
	// 		if line != prevline {
	// 			n += p.newline()
	// 		}

	// 		if line-prevline > 1 {
	// 			flush()
	// 			n += p.newline()
	// 		}
	// 	}

	// 	switch node := item.(type) {
	// 	case *ast.Comment:
	// 		if prevline != line {
	// 			flush()
	// 			if line-prevline == 1 {
	// 				n += p.newline()
	// 			}
	// 		}

	// 		n += p.comment(node)

	// 		if prevline == line {
	// 			n += p.write("\t")
	// 		}
	// 	case *ast.Verb:
	// 		if tabs == nil {
	// 			tabs = tabwriter.NewWriter(buffer, 0, 2, 1, ' ', 0)
	// 			p.writer = tabs
	// 		}

	// 		n += p.verbHead(node)

	// 		if node.Body != nil {
	// 			if !node.Body.IsEmpty() {
	// 				n += p.write("\t{\t\n")
	// 				flush()
	// 				p.indent()
	// 				n += p.verbBody(node.Body)
	// 				p.outdent()
	// 				n += p.write("}")
	// 			} else {
	// 				n += p.write("\t{}\t")
	// 			}
	// 		} else {
	// 			n += p.write("\t\t")
	// 		}
	// 	}

	// 	prevline = line
	// }

	// n += p.newline()
	// flush()
	// return
}

func (p *Printer) comment(c *ast.Comment) int {
	return p.write(c.Text)
}

func (p *Printer) verbHead(verb *ast.Verb) (n int) {
	n += p.identifier(verb.Key) + p.write("\t")
	n += p.valueGroup(verb.Values, func(n int) string {
		if n == 0 {
			return ""
		}

		return " "
	})

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
		n += p.write(pre(i))

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
	if met.Arguments != nil {
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
