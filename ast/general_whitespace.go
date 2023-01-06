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
