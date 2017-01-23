package parser

import (
	"errors"
	"fmt"

	"github.com/tyler-johnson/forge/forge/fcl/ast"
	"github.com/tyler-johnson/forge/forge/fcl/scanner"
	"github.com/tyler-johnson/forge/forge/fcl/token"
)

type Parser struct {
	Source []byte
	sc     *scanner.Scanner

	tok  token.Token
	n    int // max = 1
	line int
}

func New(src []byte) *Parser {
	return &Parser{
		Source: src,
	}
}

func (p *Parser) Parse() (*ast.File, error) {
	p.Reset()
	f := &ast.File{}
	startpos := p.peek()

	n, err := p.verbBody()
	if err != nil {
		return nil, err
	}

	tok := p.scan()
	if tok.Type != token.EOF {
		return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
	}

	n.StartPos = startpos.Pos
	n.EndPos = tok.Pos
	f.Root = n
	return f, nil
}

func (p *Parser) Reset() {
	p.sc = scanner.New(p.Source)
	p.n = 0
	p.line = 0
}

func (p *Parser) scan() token.Token {
	if p.n != 0 {
		p.n = 0
		return p.tok
	}

	p.tok = p.sc.Scan()
	return p.tok
}

func (p *Parser) unscan() {
	p.n = 1
}

func (p *Parser) peek() token.Token {
	tok := p.scan()
	p.unscan()
	return tok
}

func (p *Parser) verbBody() (*ast.VerbBody, error) {
	b := &ast.VerbBody{}

	if n := p.comment(); n != nil {
		b.Add(n)
	}

	for {
		tok := p.peek()
		prevline := p.line
		p.line = tok.Pos.Line
		if p.line <= prevline {
			if len(b.Items) > 0 {
				return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
			} else {
				break
			}
		}

		addedNode := false
		v, err := p.verb()
		if err != nil {
			return nil, err
		}

		if v != nil {
			addedNode = true
			b.Add(v)
		}

		if n := p.comment(); n != nil {
			addedNode = true
			b.Add(n)
		}

		if !addedNode {
			break
		}
	}

	return b, nil
}

func (p *Parser) comment() *ast.Comment {
	tok := p.scan()
	if tok.Type != token.COMMENT {
		p.unscan()
		return nil
	}

	c := ast.NewComment(tok.Pos)
	c.Text = tok.Text
	return c
}

func (p *Parser) verb() (*ast.Verb, error) {
	// generate key
	id := &ast.Identifier{
		RawPos: p.peek().Pos,
	}

	// get modifiers
	mods := p.modifiers()
	id.SetModifier(mods...)

	// look for standard identifier
	tok := p.scan()
	if tok.Type != token.IDENT || tok.Pos.Line != p.line {
		// we expect immediate identifier if there are modifiers
		if len(mods) > 0 {
			return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
		} else {
			p.unscan()
			return nil, nil
		}
	}
	id.Value = tok.Text

	// parse for values
	values, err := p.valueGroup()
	if err != nil {
		return nil, err
	}

	// create the verb
	verb := &ast.Verb{}
	verb.Key = id
	verb.Values = values

	// parse for body
	tok = p.scan()
	if tok.Type == token.LBRACE && tok.Pos.Line == p.line {
		b, err := p.verbBody()
		if err != nil {
			return nil, err
		}

		b.StartPos = tok.Pos
		tok = p.scan()
		if tok.Type != token.RBRACE || tok.Pos.Line != p.line {
			return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
		}

		b.EndPos = tok.Pos
		verb.Body = b
	} else {
		p.unscan()
	}

	return verb, nil
}

func (p *Parser) modifiers() []string {
	mods := make([]string, 0)

	for {
		tok := p.scan()
		if (tok.Type == token.AT || tok.Type == token.BANG) && tok.Pos.Line == p.line {
			mods = append(mods, tok.Text)
		} else {
			p.unscan()
			break
		}
	}

	return mods
}

func (p *Parser) valueGroup() (*ast.ValueGroup, error) {
	vg := &ast.ValueGroup{}

	for {
		value, err := p.value()
		if err != nil {
			return nil, err
		}

		if value != nil {
			vg.Add(value)
		} else {
			break
		}
	}

	return vg, nil
}

func (p *Parser) value() (ast.Node, error) {
	// check for literal
	if lit := p.literal(); lit != nil {
		return lit, nil
	}

	// generate key and get modifiers
	id := &ast.Identifier{
		RawPos:    p.peek().Pos,
		Modifiers: p.modifiers(),
	}

	// test for $ which means variable
	tok := p.scan()
	variable := false
	if tok.Type == token.DOLLAR && tok.Pos.Line == p.line {
		variable = true
	} else {
		p.unscan()
	}

	// next look for identity
	tok = p.scan()
	if tok.Type != token.IDENT || tok.Pos.Line != p.line {
		// if there are modifiers or is a variable we expect an identifier
		if len(id.Modifiers) > 0 || variable {
			return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
		} else {
			p.unscan()
			return nil, nil
		}
	}
	id.Value = tok.Text

	// parse paths
	path, err := p.path()
	if err != nil {
		return nil, err
	}

	// process as variable or method call
	if variable {
		return &ast.Variable{
			Key:  id,
			Path: path,
		}, nil
	} else {
		method := &ast.MethodCall{
			Key:  id,
			Path: path,
		}

		paren := p.scan()
		if paren.Type == token.LPAREN && paren.Pos.Line == p.line {
			args, err := p.valueGroup()
			if err != nil {
				return nil, err
			}

			close := p.scan()
			if close.Type != token.RPAREN || close.Pos.Line != p.line {
				return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
			}

			method.Arguments = args
		} else {
			p.unscan()
		}

		return method, nil
	}
}

func (p *Parser) literal() *ast.Literal {
	tok := p.scan()
	if tok.Pos.Line != p.line {
		p.unscan()
		return nil
	}

	l := &ast.Literal{
		RawPos: tok.Pos,
	}

	switch tok.Type {
	case token.STRING, token.HEREDOC:
		l.Type = ast.STRING
	case token.NUMBER:
		l.Type = ast.NUMBER
	case token.BOOL:
		l.Type = ast.BOOL
	case token.NULL:
		l.Type = ast.NULL
	default:
		p.unscan()
		return nil
	}

	l.Value = tok.Value()
	return l
}

func (p *Parser) path() (*ast.Path, error) {
	path := &ast.Path{}

loop:
	for {
		sep := p.scan()
		if sep.Pos.Line != p.line {
			p.unscan()
			break
		}

		switch sep.Type {
		case token.LBRACK:
			value, err := p.value()
			if err != nil {
				return nil, err
			}

			if value == nil {
				return nil, errors.New(fmt.Sprintf("Unexpected Token %s", p.peek()))
			}

			next := p.scan()
			if next.Type != token.RBRACK && next.Pos.Line == p.line {
				return nil, errors.New(fmt.Sprintf("Unexpected Token %s", p.peek()))
			}

			path.Add(&ast.PathPart{
				RawPos: value.Pos(),
				Type:   ast.COMPOUND,
				Value:  value,
			})
		case token.PERIOD:
			tok := p.scan()
			if tok.Type != token.IDENT && tok.Pos.Line == p.line {
				return nil, errors.New(fmt.Sprintf("Unexpected Token %s", tok))
			}

			path.Add(&ast.PathPart{
				RawPos: tok.Pos,
				Type:   ast.SIMPLE,
				Value:  tok.Text,
			})
		default:
			p.unscan()
			break loop
		}
	}

	return path, nil
}
