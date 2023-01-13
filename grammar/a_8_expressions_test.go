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

func TestNumber(t *testing.T) {
	testCases := []testCase[ast.Number]{
		{in: "123", want: &ast.UnsignedNumber{Value: 123}},
		{in: "32'd42", want: &ast.DecimalNumberUnsigned{Size: 32, Value: 42}},
		{in: "32'b101", want: &ast.BinaryNumber{Value: ast.MaskedInt{Size: 32, Base: 2, V: 5}}},
		{in: "32'hF0", want: &ast.HexNumber{Value: ast.MaskedInt{Size: 32, Base: 16, V: 240}}},
		{in: "32'o77", want: &ast.OctalNumber{Value: ast.MaskedInt{Size: 32, Base: 8, V: 63}}},
		{in: "32'd100_002", want: &ast.DecimalNumberUnsigned{Size: 32, Value: 100002}},
		{in: "123.456", want: &ast.FixedPointNumber{Value: 123.456}},
		{in: "123_456.789", want: &ast.FixedPointNumber{Value: 123456.789}},
		{in: "123.456e7", want: &ast.FloatingPointNumber{Value: 123.456e7}},
		{in: "123.456e-2", want: &ast.FloatingPointNumber{Value: 123.456e-2}},
		{in: "32'hFX0F8ZZ2", want: &ast.HexNumber{Value: ast.MaskedInt{Size: 32, Base: 16, V: 0xF00F8002, X: 0x0F000000, Z: 0x00000FF0}}},
		{in: "\t32 'h F0", want: &ast.HexNumber{Value: ast.MaskedInt{Size: 32, Base: 16, V: 240}}},
	}

	for _, tc := range testCases {
		validate(t, "Number", Number, tc)
	}
}

func TestUnbasedUnsizedLiteral(t *testing.T) {
	testCases := []testCase[*ast.UnbasedUnsizedLiteral]{
		{in: "'0", want: &ast.UnbasedUnsizedLiteral{Value: '0'}},
		{in: "'1", want: &ast.UnbasedUnsizedLiteral{Value: '1'}},
		{in: "'x", want: &ast.UnbasedUnsizedLiteral{Value: 'x'}},
		{in: "'z", want: &ast.UnbasedUnsizedLiteral{Value: 'z'}},
	}

	for _, tc := range testCases {
		validate(t, "UnbasedUnsizedLiteral", UnbasedUnsizedLiteral, tc)
	}
}

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

func TestString(t *testing.T) {
	testCases := []testCase[*ast.StringLiteral]{
		{in: `""`, want: &ast.StringLiteral{Text: ""}},
		{in: `"hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `  "hello"`, want: &ast.StringLiteral{Text: "hello"}},
		{in: `"hello\nworld"`, want: &ast.StringLiteral{Text: "hello\nworld"}},
		{in: `"hello\n\"world\\\"world"`, want: &ast.StringLiteral{Text: "hello\n\"world\\\"world"}},
	}

	for _, tc := range testCases {
		validate(t, "StringLiteral", StringLiteral, tc)
	}
}

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
