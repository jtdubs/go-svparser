package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

func TestConstantPrimary(t *testing.T) {
	testCases := []testCase[ast.PrimaryLiteral]{
		{in: `32`, want: &ast.UnsignedNumber{Value: 32}},
		{in: `"hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `42ms`, want: &ast.TimeLiteral{Unit: &ast.TimeUnit{Op: ast.MS}, Number: &ast.UnsignedNumber{Value: 42}}},
		{in: `  38us`, want: &ast.TimeLiteral{Unit: &ast.TimeUnit{Op: ast.US}, Number: &ast.UnsignedNumber{Value: 38}}},
		{in: `'0`, want: &ast.UnbasedUnsizedLiteral{Value: '0'}},
	}

	for _, tc := range testCases {
		validate(t, "ConstantPrimary", ConstantPrimary, tc)
	}
}

func TestPrimaryLiteral(t *testing.T) {
	testCases := []testCase[ast.PrimaryLiteral]{
		{in: `32`, want: &ast.UnsignedNumber{Value: 32}},
		{in: `"hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `42ms`, want: &ast.TimeLiteral{Unit: &ast.TimeUnit{Op: ast.MS}, Number: &ast.UnsignedNumber{Value: 42}}},
		{in: `'0`, want: &ast.UnbasedUnsizedLiteral{Value: '0'}},
	}

	for _, tc := range testCases {
		validate(t, "PrimaryLiteral", PrimaryLiteral, tc)
	}
}

func TestConstantBitSelect(t *testing.T) {
	testCases := []testCase[*ast.ConstantBitSelect]{
		{
			in: ` [42]`,
			want: &ast.ConstantBitSelect{
				Exprs: []ast.ConstantExpression{
					&ast.UnsignedNumber{Value: 42},
				},
			},
		},
		{
			in: ` [ 42  ]  [ 3+  7 ]`,
			want: &ast.ConstantBitSelect{
				Exprs: []ast.ConstantExpression{
					&ast.UnsignedNumber{Value: 42},
					&ast.ConstantBinaryExpression{
						Left:  &ast.UnsignedNumber{Value: 3},
						Op:    &ast.BinaryOperator{Op: ast.BinaryAdd},
						Right: &ast.UnsignedNumber{Value: 7},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		validateTrace(t, "ConstantBitSelect", ConstantBitSelect, tc)
	}
}
