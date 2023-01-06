package ast

import (
	"fmt"
)

type Whitespace interface {
	isWhitespace()
}

type Spaces struct {
	Token
	Text string
}

func (c *Spaces) String() string {
	return fmt.Sprintf("Spaces(%q)", c.Text)
}

func (c *Spaces) Bake() error {
	c.Text = c.Token.Value()
	return nil
}

func (*Spaces) isWhitespace() {}

type Comment interface {
	isComment()
}

type OneLineComment struct {
	Token
	StartT, EndT, TextT Token
	Text                string
}

func (c *OneLineComment) String() string {
	return fmt.Sprintf("OneLineComment(%q)", c.Text)
}

func (c *OneLineComment) Bake() error {
	c.Text = c.TextT.Value()
	return nil
}

func (*OneLineComment) isComment()    {}
func (*OneLineComment) isWhitespace() {}

type BlockComment struct {
	Token
	StartT, EndT, TextT Token
	Text                string
}

func (c *BlockComment) String() string {
	return fmt.Sprintf("BlockComment(%q)", c.Text)
}

func (c *BlockComment) Bake() error {
	c.Text = c.TextT.Value()
	return nil
}

func (*BlockComment) isComment()    {}
func (*BlockComment) isWhitespace() {}
