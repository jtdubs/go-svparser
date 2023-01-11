package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type AttributeInstance struct {
	nom.Span[rune]
	Specs []*AttrSpec
}

type AttrSpec struct {
	nom.Span[rune]
	Name *AttrName
	Expr ConstantExpression
}

type AttrName struct {
	nom.Span[rune]
	ID Identifier
}

func (i *AttrName) String() string {
	return fmt.Sprintf("AttrName(%v)", i.ID)
}
