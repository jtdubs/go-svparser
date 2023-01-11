package grammar

import (
	"testing"

	"github.com/jtdubs/go-svparser/ast"
)

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
