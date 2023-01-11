package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type ConstantExpression interface {
	isConstantExpression()
}

type ConstantUnaryExpression struct {
	nom.Span[rune]
	Op      *UnaryOperator
	Attrs   []*AttributeInstance
	Primary ConstantPrimary
}

func (e *ConstantUnaryExpression) String() string {
	return fmt.Sprintf("ConstantUnaryExpression(%v, %v)", e.Op, e.Primary)
}

func (*ConstantUnaryExpression) isConstantExpression() {}

type ConstantBinaryExpression struct {
	nom.Span[rune]
	Op          *BinaryOperator
	Attrs       []*AttributeInstance
	Left, Right ConstantExpression
}

func (e *ConstantBinaryExpression) String() string {
	return fmt.Sprintf("ConstantBinaryExpression(%v, %v, %v)", e.Left, e.Op, e.Right)
}

func (*ConstantBinaryExpression) isConstantExpression() {}

type ConstantTernaryExpression struct {
	nom.Span[rune]
	Attrs          []*AttributeInstance
	Cond, If, Else ConstantExpression
}

func (e *ConstantTernaryExpression) String() string {
	return fmt.Sprintf("ConstantTernaryExpression(%v, %v, %v)", e.Cond, e.If, e.Else)
}

func (*ConstantTernaryExpression) isConstantExpression() {}
