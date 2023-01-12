package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type Comment interface {
	isComment()
}

type OneLineComment struct {
	Token
	TextT nom.Span[rune]
	Text  string
}

func (c *OneLineComment) String() string {
	return fmt.Sprintf("OneLineComment(%q)", c.Text)
}

func (c *OneLineComment) Bake() error {
	c.Text = string(c.TextT.Value())
	return nil
}

func (*OneLineComment) isComment()    {}
func (*OneLineComment) isWhitespace() {}

type BlockComment struct {
	Token
	TextT nom.Span[rune]
	Text  string
}

func (c *BlockComment) String() string {
	return fmt.Sprintf("BlockComment(%q)", c.Text)
}

func (c *BlockComment) Bake() error {
	c.Text = string(c.TextT.Value())
	return nil
}

func (*BlockComment) isComment()    {}
func (*BlockComment) isWhitespace() {}
