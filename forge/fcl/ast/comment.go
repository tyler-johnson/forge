package ast

import "github.com/tyler-johnson/forge/forge/fcl/token"

type Comment struct {
	pos  token.Pos
	Text string
}

func NewComment(pos token.Pos) *Comment {
	return &Comment{pos: pos}
}

func (f *Comment) NodeName() string {
	return "Comment"
}

func (c *Comment) Pos() token.Pos {
	return c.pos
}
