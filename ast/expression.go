package ast

import (
	"fmt"
)

type ConstantExpression interface {
	isConstantExpression()
}

type ConstantUnaryExpression struct {
	Token
	Op      *UnaryOperator
	Attrs   []*AttributeInstance
	Primary ConstantPrimary
}

func (e *ConstantUnaryExpression) String() string {
	return fmt.Sprintf("ConstantUnaryExpression(%v, %v)", e.Op, e.Primary)
}

func (*ConstantUnaryExpression) isConstantExpression() {}

type ConstantBinaryExpression struct {
	Token
	Op          *BinaryOperator
	Attrs       []*AttributeInstance
	Left, Right ConstantExpression
}

func (e *ConstantBinaryExpression) String() string {
	return fmt.Sprintf("ConstantBinaryExpression(%v, %v, %v)", e.Left, e.Op, e.Right)
}

func (*ConstantBinaryExpression) isConstantExpression() {}

type ConstantTernaryExpression struct {
	Token
	Attrs          []*AttributeInstance
	Cond, If, Else ConstantExpression
}

func (e *ConstantTernaryExpression) String() string {
	return fmt.Sprintf("ConstantTernaryExpression(%v, %v, %v)", e.Cond, e.If, e.Else)
}

func (*ConstantTernaryExpression) isConstantExpression() {}

type ConstantBitSelect struct {
	Token
	Exprs []ConstantExpression
}

func (e *ConstantBitSelect) String() string {
	return fmt.Sprintf("ConstantBitSelect(%v)", e.Exprs)
}
