package ast

import (
	"fmt"

	"github.com/jtdubs/go-nom"
)

type UnaryOperatorType int

const (
	UnaryPlus UnaryOperatorType = iota
	UnaryMinus
	UnaryLogicalNot
	UnaryBinaryNot
	UnaryAnd
	UnaryNand
	UnaryOr
	UnaryNor
	UnaryXor
	UnaryXnor
)

var unaryOperatorNames = map[UnaryOperatorType]string{
	UnaryPlus:       "Plus",
	UnaryMinus:      "Minus",
	UnaryLogicalNot: "LogicalNot",
	UnaryBinaryNot:  "BinaryNot",
	UnaryAnd:        "And",
	UnaryNand:       "Nand",
	UnaryOr:         "Or",
	UnaryNor:        "Nor",
	UnaryXor:        "Xor",
	UnaryXnor:       "Xnor",
}

type UnaryOperator struct {
	nom.Span[rune]
	Op UnaryOperatorType
}

func (u *UnaryOperator) String() string {
	return fmt.Sprintf("UnaryOperator(%v)", unaryOperatorNames[u.Op])
}
