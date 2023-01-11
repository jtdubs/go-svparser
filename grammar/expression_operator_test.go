package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestUnaryOperator(t *testing.T) {
	testCases := []testCase[*ast.UnaryOperator]{
		{in: "~&", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNand}},
		{in: "~|", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNor}},
		{in: "~^", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "^~", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "+", want: &ast.UnaryOperator{Op: ast.UnaryPositive}},
		{in: "-", want: &ast.UnaryOperator{Op: ast.UnaryNegate}},
		{in: "!", want: &ast.UnaryOperator{Op: ast.UnaryLogicalNegation}},
		{in: "&", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionAnd}},
		{in: "|", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionOr}},
		{in: "^", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionXor}},
		{in: "~", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNot}},
		{in: " ~", want: &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNot}},
	}

	for _, tc := range testCases {
		validate(t, "UnaryOperator", UnaryOperator, tc)
	}
}

func TestBinaryOperator(t *testing.T) {
	testCases := []testCase[*ast.BinaryOperator]{
		{in: "<<<", want: &ast.BinaryOperator{Op: ast.BinaryArithmeticShiftLeft}},
		{in: ">>>", want: &ast.BinaryOperator{Op: ast.BinaryArithmeticShiftRight}},
		{in: "<->", want: &ast.BinaryOperator{Op: ast.BinaryLogicalIff}},
		{in: "===", want: &ast.BinaryOperator{Op: ast.BinaryCaseEquals}},
		{in: "!==", want: &ast.BinaryOperator{Op: ast.BinaryCaseNotEquals}},
		{in: "==?", want: &ast.BinaryOperator{Op: ast.BinaryWildcardEquals}},
		{in: "!=?", want: &ast.BinaryOperator{Op: ast.BinaryWildcardNotEquals}},
		{in: "->", want: &ast.BinaryOperator{Op: ast.BinaryLogicalImplies}},
		{in: "**", want: &ast.BinaryOperator{Op: ast.BinaryExp}},
		{in: "^~", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "~^", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "<<", want: &ast.BinaryOperator{Op: ast.BinaryLogicalShiftLeft}},
		{in: ">>", want: &ast.BinaryOperator{Op: ast.BinaryLogicalShiftRight}},
		{in: "&&", want: &ast.BinaryOperator{Op: ast.BinaryLogicalAnd}},
		{in: "||", want: &ast.BinaryOperator{Op: ast.BinaryLogicalOr}},
		{in: "<=", want: &ast.BinaryOperator{Op: ast.BinaryLessThanEqual}},
		{in: ">=", want: &ast.BinaryOperator{Op: ast.BinaryGreaterThanEqual}},
		{in: "!=", want: &ast.BinaryOperator{Op: ast.BinaryLogicalNotEquals}},
		{in: "==", want: &ast.BinaryOperator{Op: ast.BinaryLogicalEquals}},
		{in: "<", want: &ast.BinaryOperator{Op: ast.BinaryLessThan}},
		{in: ">", want: &ast.BinaryOperator{Op: ast.BinaryGreaterThan}},
		{in: "^", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseXor}},
		{in: "+", want: &ast.BinaryOperator{Op: ast.BinaryAdd}},
		{in: "-", want: &ast.BinaryOperator{Op: ast.BinarySubtract}},
		{in: "*", want: &ast.BinaryOperator{Op: ast.BinaryMultiply}},
		{in: "/", want: &ast.BinaryOperator{Op: ast.BinaryDivide}},
		{in: "%", want: &ast.BinaryOperator{Op: ast.BinaryModulus}},
		{in: "&", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseAnd}},
		{in: "|", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseOr}},
		{in: "\t|", want: &ast.BinaryOperator{Op: ast.BinaryBitwiseOr}},
	}

	for _, tc := range testCases {
		validate(t, "BinaryOperator", BinaryOperator, tc)
	}
}

func TestModulePathUnaryOperator(t *testing.T) {
	testCases := []testCase[*ast.UnaryModulePathOperator]{
		{in: "~&", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionNand}},
		{in: "~|", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionNor}},
		{in: "~^", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "^~", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionXnor}},
		{in: "!", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalNegation}},
		{in: "&", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionAnd}},
		{in: "|", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionOr}},
		{in: "^", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionXor}},
		{in: "~", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionNot}},
		{in: "\t~", want: &ast.UnaryModulePathOperator{Op: ast.UnaryLogicalReductionNot}},
	}

	for _, tc := range testCases {
		validate(t, "UnaryModulePathOperator", UnaryModulePathOperator, tc)
	}
}

func TestBinaryModulePathOperator(t *testing.T) {
	testCases := []testCase[*ast.BinaryModulePathOperator]{
		{in: "^~", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "~^", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseXnor}},
		{in: "&&", want: &ast.BinaryModulePathOperator{Op: ast.BinaryLogicalAnd}},
		{in: "||", want: &ast.BinaryModulePathOperator{Op: ast.BinaryLogicalOr}},
		{in: "!=", want: &ast.BinaryModulePathOperator{Op: ast.BinaryLogicalNotEquals}},
		{in: "==", want: &ast.BinaryModulePathOperator{Op: ast.BinaryLogicalEquals}},
		{in: "^", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseXor}},
		{in: "&", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseAnd}},
		{in: "|", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseOr}},
		{in: "\t|", want: &ast.BinaryModulePathOperator{Op: ast.BinaryBitwiseOr}},
	}

	for _, tc := range testCases {
		validate(t, "BinaryModulePathOperator", BinaryModulePathOperator, tc)
	}
}
