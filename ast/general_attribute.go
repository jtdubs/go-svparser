package ast

import (
	"fmt"
)

type AttributeInstance struct {
	Token
	Specs []*AttrSpec
}

type AttrSpec struct {
	Token
	Name *AttrName
	Expr ConstantExpression
}

type AttrName struct {
	Token
	ID Identifier
}

func (i *AttrName) String() string {
	return fmt.Sprintf("AttrName(%v)", i.ID)
}
