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

type CIdentifier struct {
	nom.Span[rune]
	Name string
}

func (i *CIdentifier) String() string {
	return fmt.Sprintf("CIdentifier(%v)", i.Name)
}

func (i *CIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return nil
}

type SystemTfIdentifier struct {
	nom.Span[rune]
	Name string
}

func (i *SystemTfIdentifier) String() string {
	return fmt.Sprintf("SystemTfIdentifier(%v)", i.Name)
}

func (i *SystemTfIdentifier) Bake() error {
	i.Name = string(i.Span.Value())
	return nil
}

type TaskIdentifier struct {
	nom.Span[rune]
	ID Identifier
}

func (i *TaskIdentifier) String() string {
	return fmt.Sprintf("TaskIdentifier(%v)", i.ID)
}
