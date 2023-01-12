package ast

import (
	"fmt"
)

type UnaryOperatorType int

const (
	UnaryPositive UnaryOperatorType = iota
	UnaryNegate
	UnaryLogicalNegation
	UnaryLogicalReductionNot
	UnaryLogicalReductionAnd
	UnaryLogicalReductionNand
	UnaryLogicalReductionOr
	UnaryLogicalReductionNor
	UnaryLogicalReductionXor
	UnaryLogicalReductionXnor
)

var unaryOperatorNames = map[UnaryOperatorType]string{
	UnaryPositive:             "Positive",
	UnaryNegate:               "Negate",
	UnaryLogicalNegation:      "LogicalNegation",
	UnaryLogicalReductionNot:  "LogicalReductionNot",
	UnaryLogicalReductionAnd:  "LogicalReductionAnd",
	UnaryLogicalReductionNand: "LogicalReductionNand",
	UnaryLogicalReductionOr:   "LogicalReductionOr",
	UnaryLogicalReductionNor:  "LogicalReductionNor",
	UnaryLogicalReductionXor:  "LogicalReductionXor",
	UnaryLogicalReductionXnor: "LogicalReductionXnor",
}

type UnaryOperator struct {
	Token
	Op UnaryOperatorType
}

func (u *UnaryOperator) String() string {
	return fmt.Sprintf("UnaryOperator(%v)", unaryOperatorNames[u.Op])
}

type BinaryOperatorType int

const (
	BinaryAdd BinaryOperatorType = iota
	BinarySubtract
	BinaryMultiply
	BinaryDivide
	BinaryModulus
	BinaryExp
	BinaryBitwiseAnd
	BinaryBitwiseOr
	BinaryBitwiseXor
	BinaryBitwiseXnor
	BinaryLogicalShiftLeft
	BinaryLogicalShiftRight
	BinaryArithmeticShiftLeft
	BinaryArithmeticShiftRight
	BinaryLogicalAnd
	BinaryLogicalOr
	BinaryLogicalImplies
	BinaryLogicalIff
	BinaryLessThan
	BinaryLessThanEqual
	BinaryGreaterThan
	BinaryGreaterThanEqual
	BinaryCaseEquals
	BinaryCaseNotEquals
	BinaryLogicalEquals
	BinaryLogicalNotEquals
	BinaryWildcardEquals
	BinaryWildcardNotEquals
)

var binaryOperatorNames = map[BinaryOperatorType]string{
	BinaryAdd:                  "Add",
	BinarySubtract:             "Subtract",
	BinaryMultiply:             "Multiply",
	BinaryDivide:               "Divide",
	BinaryModulus:              "Modulus",
	BinaryExp:                  "Exp",
	BinaryBitwiseAnd:           "BitwiseAnd",
	BinaryBitwiseOr:            "BitwiseOr",
	BinaryBitwiseXor:           "BitwiseXor",
	BinaryBitwiseXnor:          "BitwiseXnor",
	BinaryLogicalShiftLeft:     "LogicalShiftLeft",
	BinaryLogicalShiftRight:    "LogicalShiftRight",
	BinaryArithmeticShiftLeft:  "ArithmeticShiftLeft",
	BinaryArithmeticShiftRight: "ArithmeticShiftRight",
	BinaryLogicalAnd:           "LogicalAnd",
	BinaryLogicalOr:            "LogicalOr",
	BinaryLogicalImplies:       "LogicalImplies",
	BinaryLogicalIff:           "LogicalIff",
	BinaryLessThan:             "LessThan",
	BinaryLessThanEqual:        "LessThanEqual",
	BinaryGreaterThan:          "GreaterThan",
	BinaryGreaterThanEqual:     "GreaterThanEqual",
	BinaryCaseEquals:           "CaseEquals",
	BinaryCaseNotEquals:        "CaseNotEquals",
	BinaryLogicalEquals:        "LogicalEquals",
	BinaryLogicalNotEquals:     "LogicalNotEquals",
	BinaryWildcardEquals:       "WildcardEquals",
	BinaryWildcardNotEquals:    "WildcardNotEquals",
}

type BinaryOperator struct {
	Token
	Op BinaryOperatorType
}

func (b *BinaryOperator) String() string {
	return fmt.Sprintf("BinaryOperator(%v)", binaryOperatorNames[b.Op])
}

type IncOrDecOperatorType int

const (
	Inc IncOrDecOperatorType = iota
	Dec
)

var incOrDecOperatorNames = map[IncOrDecOperatorType]string{
	Inc: "Inc",
	Dec: "Dec",
}

type IncOrDecOperator struct {
	Token
	Op IncOrDecOperatorType
}

func (b *IncOrDecOperator) String() string {
	return fmt.Sprintf("IncOrDecOperator(%v)", incOrDecOperatorNames[b.Op])
}

type UnaryModulePathOperator struct {
	Token
	Op UnaryOperatorType
}

func (b *UnaryModulePathOperator) String() string {
	return fmt.Sprintf("UnaryModulePathOperator(%v)", unaryOperatorNames[b.Op])
}

type BinaryModulePathOperator struct {
	Token
	Op BinaryOperatorType
}

func (b *BinaryModulePathOperator) String() string {
	return fmt.Sprintf("BinaryModulePathOperator(%v)", binaryOperatorNames[b.Op])
}
