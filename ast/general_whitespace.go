package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type Whitespace interface {
	isWhitespace()
}

type Spaces struct {
	nom.Span[rune]
	Text string
}

func (c *Spaces) String() string {
	return fmt.Sprintf("Spaces(%q)", c.Text)
}

func (c *Spaces) Bake() error {
	c.Text = string(c.Span.Value())
	return nil
}

func (*Spaces) isWhitespace() {}
