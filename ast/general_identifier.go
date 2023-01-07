package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type Identifier interface {
	isIdentifier()
}

type SimpleIdentifier struct {
	nom.Span[rune]
	Name string
}

func (i *SimpleIdentifier) String() string {
	return fmt.Sprintf("SimpleIdentifier(%v)", i.Name)
}

func (i *SimpleIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return nil
}

func (*SimpleIdentifier) isIdentifier() {}

type EscapedIdentifier struct {
	nom.Span[rune]
	SlashT, NameT nom.Span[rune]
	Name          string
}

func (i *EscapedIdentifier) String() string {
	return fmt.Sprintf("EscapedIdentifier(%v)", i.Name)
}

func (i *EscapedIdentifier) Bake() error {
	i.Name = string(i.NameT.Value())
	return nil
}

func (*EscapedIdentifier) isIdentifier() {}
