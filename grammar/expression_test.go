package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestConstantExpression(t *testing.T) {
	testCases := []testCase[ast.ConstantExpression]{
		{
			in:   "42",
			want: &ast.UnsignedNumber{Value: 42},
		},
		{
			in: "~10",
			want: &ast.ConstantUnaryExpression{
				Op:      &ast.UnaryOperator{Op: ast.UnaryLogicalReductionNot},
				Primary: &ast.UnsignedNumber{Value: 10},
			},
		},
		{
			in: "\t10 + 20",
			want: &ast.ConstantBinaryExpression{
				Op:    &ast.BinaryOperator{Op: ast.BinaryAdd},
				Left:  &ast.UnsignedNumber{Value: 10},
				Right: &ast.UnsignedNumber{Value: 20},
			},
		},
	}

	for _, tc := range testCases {
		validate(t, "ConstantExpression", ConstantExpression, tc)
	}
}
